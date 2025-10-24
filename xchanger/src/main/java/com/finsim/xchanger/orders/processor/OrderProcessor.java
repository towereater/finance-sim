package com.finsim.xchanger.orders.processor;

import java.time.Instant;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;
import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

import com.finsim.xchanger.banks.model.Bank;
import com.finsim.xchanger.banks.service.BankService;
import com.finsim.xchanger.common.model.Price;
import com.finsim.xchanger.dossiers.model.Dossier;
import com.finsim.xchanger.dossiers.model.DossierStock;
import com.finsim.xchanger.dossiers.service.DossierService;
import com.finsim.xchanger.orders.model.Order;
import com.finsim.xchanger.orders.model.OrderTransaction;
import com.finsim.xchanger.orders.service.OrderService;
import com.finsim.xchanger.payments.model.InsertPaymentRequest;
import com.finsim.xchanger.payments.model.Payment;
import com.finsim.xchanger.payments.service.PaymentService;
import com.finsim.xchanger.stocks.model.Stock;
import com.finsim.xchanger.stocks.model.StockPrices;
import com.finsim.xchanger.stocks.service.StockService;

@Component
public class OrderProcessor {
    @Autowired
    private BankService bankService;

    @Autowired
    private StockService stockService;

    @Autowired
    private DossierService dossierService;

    @Autowired
    private OrderService orderService;

    @Autowired
    private PaymentService paymentService;

    @KafkaListener(topics = "${xchanger.kafka.topic.orders}")
    public void elaborateOrder(String id) {
        System.out.printf("Elaboration of order %s started\n", id);

        // Get order data from given id
        Optional<Order> orderOptional = orderService.findOrderById(id);
        if (orderOptional.isEmpty()) {
            System.out.printf("No order with id %s found\n", id);
            return;
        }
        Order order = orderOptional.get();
        
        // Elaborate order depending on order type
        switch (order.getType()) {
            case "BUY":
                elaborateBuyOrder(order);
                break;
            case "SELL":
                elaborateSellOrder(order);
                break;
            default:
                System.out.printf("Order type %s unknown\n", order.getType());
                break;
        }

        System.out.printf("Elaboration of order %s completed\n", id);
    }

    public void elaborateBuyOrder(Order buyOrder) {
        while (buyOrder.getLeftQuantity() > 0) {
            // Get best sell orders to start transactions with
            Optional<Order> sellOrderOptional = orderService.findBestSellOrderWithPriceLowerThan(
                buyOrder.getIsin(), buyOrder.getPrice().getAmount());
            if (sellOrderOptional.isEmpty()) {
                System.out.printf("No available sell orders with isin %s found\n", buyOrder.getIsin());
                return;
            }

            // Check whether buy and sell dossiers match
            Order sellOrder = sellOrderOptional.get();
            if (buyOrder.getDossier() == sellOrder.getDossier()) {
                System.out.printf("Buy and sell orders have matching dossiers\n");
                return;
            }

            // Execute transaction
            elaborateOrders(buyOrder, sellOrder);

            // Update buy order object
            Optional<Order> buyOrderOptional = orderService.findOrderById(buyOrder.getId());
            if (buyOrderOptional.isEmpty()) {
                System.out.printf("No buy order with id %s found\n", buyOrder.getId());
                return;
            }
            buyOrder = buyOrderOptional.get();
        }
    }

    public void elaborateSellOrder(Order sellOrder) {
        while (sellOrder.getLeftQuantity() > 0) {
            // Get best buy orders to start transactions with
            Optional<Order> buyOrderOptional = orderService.findBestBuyOrderWithPriceGreaterThan(
                sellOrder.getIsin(), sellOrder.getPrice().getAmount());
            if (buyOrderOptional.isEmpty()) {
                System.out.printf("No available sell orders with isin %s found\n", sellOrder.getIsin());
                return;
            }

            // Check whether buy and sell dossiers match
            Order buyOrder = buyOrderOptional.get();
            if (buyOrder.getDossier() == sellOrder.getDossier()) {
                System.out.printf("Buy and sell orders have matching dossiers\n");
                return;
            }

            // Execute transaction
            elaborateOrders(buyOrder, sellOrder);

            // Update sell order object
            Optional<Order> sellOrderOptional = orderService.findOrderById(sellOrder.getId());
            if (sellOrderOptional.isEmpty()) {
                System.out.printf("No sell order with id %s found\n", sellOrder.getId());
                return;
            }
            sellOrder = sellOrderOptional.get();
        }
    }

    public void elaborateOrders(Order buyOrder, Order sellOrder) {
        // Get dossier associated with buy order
        Optional<Dossier> buyDossierOptional = dossierService.findDossierById(buyOrder.getDossier());
        if (buyDossierOptional.isEmpty()) {
            System.out.printf("No dossiers with id %s found\n", buyOrder.getDossier());
            return;
        }
        Dossier buyDossier = buyDossierOptional.get();

        // Get bank associated with buy order
        Optional<Bank> bankOptional = bankService.findBankByAbi(buyDossier.getAbi());
        if (bankOptional.isEmpty()) {
            System.out.printf("No banks with abi %s found\n", buyDossier.getAbi());
            return;
        }
        Bank bank = bankOptional.get();

        // Get dossier associated with sell order
        Optional<Dossier> sellDossierOptional = dossierService.findDossierById(sellOrder.getDossier());
        if (sellDossierOptional.isEmpty()) {
            System.out.printf("No dossiers with id %s found\n", sellOrder.getDossier());
            return;
        }
        Dossier sellDossier = sellDossierOptional.get();

        // Get quantity of stocks to move between dossiers
        String isin = buyOrder.getIsin();
        int quantity = buyOrder.getLeftQuantity() >= sellOrder.getLeftQuantity()
            ? sellOrder.getLeftQuantity() : buyOrder.getLeftQuantity();

        // Create payment request before applying stocks transactions
        InsertPaymentRequest paymentRequest = paymentService.createBankTransfer(
            sellOrder.getPrice().multiplyBy(quantity).toPaymentValue(),
            buyDossier.getIban(),
            sellDossier.getName() + ' ' + sellDossier.getSurname(),
            sellDossier.getIban(),
            String.format("Buy %d %s stocks for %s",
                quantity,
                isin,
                sellOrder.getPrice().toString()
            )
        );
        ResponseEntity<Payment> res = paymentService.insertPayment(bank.getExternalApiToken(), paymentRequest);
        if (!res.getStatusCode().equals(HttpStatus.CREATED)) {
            System.out.printf("Payment request received status code %d\n", res.getStatusCode().value());
            return;
        }

        // Update sell dossier stocks
        List<DossierStock> sellDossierStocks = sellDossier.getStocks();
        Optional<DossierStock> sellDossierStockOptional = sellDossierStocks.stream()
            .filter(ds -> ds.getIsin().equals(isin))
            .findFirst();
        if (sellDossierStockOptional.isEmpty()) {
            System.out.printf("Sell dossier %s has no available stocks with isin %s\n",
                sellDossier.getId(), isin);
            return;
        }

        DossierStock sellDossierStock = sellDossierStockOptional.get();
        sellDossierStock.setTotal(sellDossierStock.getTotal() - quantity);
        if (sellDossierStock.getTotal() == 0) {
            sellDossierStocks.removeIf(ds -> ds.getIsin().equals(isin));
        } else {
            sellDossierStocks.replaceAll(ds -> ds.getIsin().equals(isin) ? sellDossierStock : ds);
        }

        sellDossier.setStocks(sellDossierStocks);
        sellDossier = dossierService.updateDossier(sellDossier);

        // Update buy dossier stocks
        List<DossierStock> buyDossierStocks = buyDossier.getStocks();

        if (buyDossierStocks == null) {
            buyDossierStocks = new ArrayList<DossierStock>();
            buyDossierStocks.add(new DossierStock(isin, quantity, quantity));
        } else {
            Optional<DossierStock> buyDossierStockOptional = buyDossierStocks.stream()
                .filter(ds -> ds.getIsin().equals(isin))
                .findFirst();

            if (buyDossierStockOptional.isEmpty()) {
                buyDossierStocks.add(new DossierStock(isin, quantity, quantity));
                buyDossierStocks.sort(new Comparator<DossierStock>() {
                    @Override
                    public int compare(DossierStock rhs, DossierStock lhs) {
                        return rhs.getIsin().compareTo(lhs.getIsin());
                    }
                });
            } else {
                DossierStock buyDossierStock = buyDossierStockOptional.get();
                buyDossierStock.setTotal(buyDossierStock.getTotal() - quantity);
                buyDossierStocks.replaceAll(ds -> ds.getIsin().equals(isin) ? buyDossierStock : ds);
            }
        }

        buyDossier.setStocks(buyDossierStocks);
        buyDossier = dossierService.updateDossier(buyDossier);

        // Update current orders
        String timestamp = Instant.now().toString();

        List<OrderTransaction> sellOrderTransactions = sellOrder.getOrderTransactions();
        if (sellOrderTransactions == null) {
            sellOrderTransactions = new ArrayList<OrderTransaction>();
        }
        sellOrderTransactions.add(new OrderTransaction(
            buyOrder.getDossier(),
            quantity,
            sellOrder.getPrice(),
            timestamp
        ));

        List<OrderTransaction> buyOrderTransactions = buyOrder.getOrderTransactions();
        if (buyOrderTransactions == null) {
            buyOrderTransactions = new ArrayList<OrderTransaction>();
        }
        buyOrderTransactions.add(new OrderTransaction(
            sellOrder.getDossier(),
            quantity,
            sellOrder.getPrice(),
            timestamp
        ));

        sellOrder.setLeftQuantity(sellOrder.getLeftQuantity() - quantity);
        sellOrder.setOrderTransactions(sellOrderTransactions);
        sellOrder = orderService.updateOrder(sellOrder);

        buyOrder.setLeftQuantity(buyOrder.getLeftQuantity() - quantity);
        buyOrder.setOrderTransactions(buyOrderTransactions);
        buyOrder = orderService.updateOrder(buyOrder);

        // Get stocks data
        Optional<Stock> stockOptional = stockService.findStockByIsin(isin);
        if (stockOptional.isEmpty()) {
            System.out.printf("No stocks with isin %s found\n", isin);
            return;
        }
        Stock stock = stockOptional.get();

        // Check stock price variation
        Boolean updated = false;
        float sellAmount = sellOrder.getPrice().getAmount();
        String sellCurrency = sellOrder.getPrice().getCurrency();

        if (stock.getPrices() == null) {
            stock.setPrices(new StockPrices(
                new Price(sellAmount, sellCurrency),
                new Price(sellAmount, sellCurrency),
                new Price(sellAmount, sellCurrency),
                new Price(sellAmount, sellCurrency)
            ));
            updated = true;
        } else {
            if (stock.getPrices().getDailyMax() == null ||
                stock.getPrices().getDailyMax().getAmount() < sellAmount) {
                stock.getPrices().setDailyMax(new Price(sellAmount, sellCurrency));
                updated = true;
            } else if (stock.getPrices().getDailyMin() == null ||
                stock.getPrices().getDailyMin().getAmount() > sellAmount) {
                stock.getPrices().setDailyMin(new Price(sellAmount, sellCurrency));
                updated = true;
            }

            if (stock.getPrices().getDailyOpening() == null) {
                stock.getPrices().setDailyOpening(new Price(sellAmount, sellCurrency));
                updated = true;
            }

            if (stock.getPrices().getDailyLast().getAmount() != sellAmount) {
                stock.getPrices().setDailyLast(new Price(sellAmount, sellCurrency));
                updated = true;
            }
        }

        // Update stock data if needed
        if (updated) {
            stock = stockService.updateStock(stock);
        }
    }
}
