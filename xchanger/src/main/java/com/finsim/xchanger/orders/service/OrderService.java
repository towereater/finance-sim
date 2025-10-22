package com.finsim.xchanger.orders.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;

import com.finsim.xchanger.banks.model.Bank;
import com.finsim.xchanger.banks.repository.BankRepository;
import com.finsim.xchanger.dossiers.model.Dossier;
import com.finsim.xchanger.dossiers.repository.DossierRepository;
import com.finsim.xchanger.orders.configuration.OrderProducer;
import com.finsim.xchanger.orders.model.InsertOrderRequest;
import com.finsim.xchanger.orders.model.Order;
import com.finsim.xchanger.orders.repository.OrderRepository;
import com.finsim.xchanger.payments.model.InsertPaymentRequest;
import com.finsim.xchanger.payments.model.Payment;
import com.finsim.xchanger.payments.service.PaymentService;

import java.util.List;
import java.util.Optional;

@Service
public class OrderService {
    @Autowired
    private OrderRepository orderRepository;

    @Autowired
    private OrderProducer orderProducer;

    @Autowired
    private DossierRepository dossierRepository;

    @Autowired
    private BankRepository bankRepository;

    @Autowired
    private PaymentService paymentService;

    public Page<Order> findAllOrders(Pageable pageable) {
        return orderRepository.findAll(pageable);
    }

    public Optional<Order> findOrderById(String id) {
        return orderRepository.findById(id);
    }

    public Order createOrder(InsertOrderRequest orderRequest) {
        Order order = new Order();
        order.setDossier(orderRequest.dossier);
        order.setIsin(orderRequest.isin);
        order.setType(orderRequest.type);
        order.setPrice(orderRequest.price);
        order.setQuantity(orderRequest.quantity);
        order.setOptions(orderRequest.options);
        order.setLeftQuantity(orderRequest.quantity);

        order = insertOrder(order);

        // Add order id to queue
        orderProducer.queueOrder(order.getId());

        return order;
    }

    public Order insertOrder(Order order) {
        return orderRepository.insert(order);
    }

    public void deleteOrder(String id) {
        orderRepository.deleteById(id);
    }

    public void elaborateOrder(String id) {
        // Get order data from given id
        Optional<Order> orderOptional = findOrderById(id);
        if (orderOptional.isEmpty()) {
            System.out.printf("No order with id %s found\n", id);
            return;
        }
        Order order = orderOptional.get();
        
        // Elaborate order depending on order type
        switch (order.getType()) {
            case "BUY":
                System.out.printf("BUY Order with id %s found\n", id);
                elaborateBuyOrder(order);
                break;
            case "SELL":
                System.out.printf("SELL Order with id %s found\n", id);
                elaborateSellOrder(order);
                break;
            default:
                System.out.printf("Order type %s unknown\n", order.getType());
                break;
        }
    }

    public void elaborateBuyOrder(Order order) {
        // Get best sell orders to start transactions with
        List<Order> sellOrders = orderRepository.findBestSellOrders(order.getIsin());
        if (sellOrders.isEmpty()) {
            System.out.printf("No available sell orders with isin %s found\n", order.getIsin());
            return;
        }
        Order sellOrder = sellOrders.getFirst();

        // Get dossier associated with starting order
        Optional<Dossier> dossierOptional = dossierRepository.findById(order.getDossier());
        if (dossierOptional.isEmpty()) {
            System.out.printf("No dossiers with id %s found\n", order.getDossier());
            return;
        }
        Dossier dossier = dossierOptional.get();

        // Get bank associated with starting order
        Optional<Bank> bankOptional = bankRepository.findByAbi(dossier.getAbi());
        if (bankOptional.isEmpty()) {
            System.out.printf("No banks with abi %s found\n", dossier.getAbi());
            return;
        }
        Bank bank = bankOptional.get();

        // Get dossier associated with sell order
        Optional<Dossier> sellDossierOptional = dossierRepository.findById(sellOrder.getDossier());
        if (sellDossierOptional.isEmpty()) {
            System.out.printf("No dossiers with id %s found\n", sellOrder.getDossier());
            return;
        }
        Dossier sellDossier = sellDossierOptional.get();

        // Elaborate current order until no matching sell orders are found
        while (order.getLeftQuantity() > 0
            && sellOrder != null
            && order.getPrice().getAmount() >= sellOrder.getPrice().getAmount()) {
            // Sell order has more stocks than needed
            if (order.getLeftQuantity() <= sellOrder.getLeftQuantity()) {
                // Create payment request before applying stocks transactions
                InsertPaymentRequest paymentRequest = paymentService.createBankTransfer(
                    sellOrder.getPrice().multiplyBy(order.getLeftQuantity()).toPaymentValue(),
                    dossier.getIban(),
                    sellDossier.getName() + ' ' + sellDossier.getSurname(),
                    sellDossier.getIban(),
                    String.format("Buy %d %s stocks for %s",
                        order.getLeftQuantity(),
                        order.getIsin(),
                        sellOrder.getPrice().toString()
                    )
                );
                ResponseEntity<Payment> res = paymentService.insertPayment(bank.externalApiToken, paymentRequest);
                if (res.getStatusCode() != HttpStatus.CREATED) {
                    System.out.printf("Payment request received status code %d\n", res.getStatusCode().value());
                    return;
                }
                System.out.printf("Payment request created successufully\n");

                // Update current orders
                sellOrder.setLeftQuantity(sellOrder.getLeftQuantity() - order.getLeftQuantity());
                orderRepository.save(sellOrder);

                order.setLeftQuantity(0);
                order = orderRepository.save(order);
            }
            // Sell order has less stocks than needed
            else {
                // Create payment request before applying stocks transactions
                InsertPaymentRequest paymentRequest = paymentService.createBankTransfer(
                    sellOrder.getPrice().multiplyBy(sellOrder.getLeftQuantity()).toPaymentValue(),
                    dossier.getIban(),
                    sellDossier.getName() + ' ' + sellDossier.getSurname(),
                    sellDossier.getIban(),
                    String.format("Buy %d %s stocks for %s",
                        sellOrder.getLeftQuantity(),
                        order.getIsin(),
                        sellOrder.getPrice().toString()
                    )
                );
                ResponseEntity<Payment> res = paymentService.insertPayment(bank.externalApiToken, paymentRequest);
                if (res.getStatusCode() != HttpStatus.CREATED) {
                    System.out.printf("Payment request received status code %d\n", res.getStatusCode().value());
                    return;
                }
                System.out.printf("Payment request created successufully\n");

                // Update current orders
                order.setLeftQuantity(order.getLeftQuantity() - sellOrder.getLeftQuantity());
                order = orderRepository.save(order);

                sellOrder.setLeftQuantity(0);
                orderRepository.save(sellOrder);

                // Get best sell orders to start transactions with
                sellOrders = orderRepository.findBestSellOrders(order.getIsin());
                if (sellOrders.isEmpty()) {
                    System.out.printf("No available sell orders with isin %s found\n", order.getIsin());
                    return;
                }
                sellOrder = sellOrders.getFirst();

                // Get dossier associated with sell order
                sellDossierOptional = dossierRepository.findById(sellOrder.getDossier());
                if (sellDossierOptional.isEmpty()) {
                    System.out.printf("No dossiers with id %s found\n", sellOrder.getDossier());
                    return;
                }
                sellDossier = sellDossierOptional.get();
            }
        }
    }

    public void elaborateSellOrder(Order order) {
        // Get best buy orders to start transactions with
        List<Order> buyOrders = orderRepository.findBestBuyOrders(order.getIsin());
        if (buyOrders.isEmpty()) {
            System.out.printf("No available sell orders with isin %s found\n", order.getIsin());
            return;
        }
        Order buyOrder = buyOrders.getFirst();

        // Get dossier associated with starting order
        Optional<Dossier> dossierOptional = dossierRepository.findById(order.getDossier());
        if (dossierOptional.isEmpty()) {
            System.out.printf("No dossiers with id %s found\n", order.getDossier());
            return;
        }
        Dossier dossier = dossierOptional.get();

        // Get bank associated with starting order
        Optional<Bank> bankOptional = bankRepository.findByAbi(dossier.getAbi());
        if (bankOptional.isEmpty()) {
            System.out.printf("No banks with abi %s found\n", dossier.getAbi());
            return;
        }
        Bank bank = bankOptional.get();

        // Get dossier associated with buy order
        Optional<Dossier> buyDossierOptional = dossierRepository.findById(buyOrder.getDossier());
        if (buyDossierOptional.isEmpty()) {
            System.out.printf("No dossiers with id %s found\n", buyOrder.getDossier());
            return;
        }
        Dossier buyDossier = buyDossierOptional.get();

        // Elaborate current order until no matching buy orders are found
        while (order.getLeftQuantity() > 0
            && buyOrder != null
            && order.getPrice().getAmount() <= buyOrder.getPrice().getAmount()) {
            // Buy order has more stocks than needed
            if (order.getLeftQuantity() <= buyOrder.getLeftQuantity()) {
                // Create payment request before applying stocks transactions
                InsertPaymentRequest paymentRequest = paymentService.createBankTransfer(
                    order.getPrice().multiplyBy(order.getLeftQuantity()).toPaymentValue(),
                    buyDossier.getIban(),
                    dossier.getName() + ' ' + dossier.getSurname(),
                    dossier.getIban(),
                    String.format("Buy %d %s stocks for %s",
                        order.getLeftQuantity(),
                        order.getIsin(),
                        order.getPrice().toString()
                    )
                );
                ResponseEntity<Payment> res = paymentService.insertPayment(bank.externalApiToken, paymentRequest);
                if (res.getStatusCode() != HttpStatus.CREATED) {
                    System.out.printf("Payment request received status code %d\n", res.getStatusCode().value());
                    return;
                }
                System.out.printf("Payment request created successufully\n");

                // Update current orders
                buyOrder.setLeftQuantity(buyOrder.getLeftQuantity() - order.getLeftQuantity());
                orderRepository.save(buyOrder);

                order.setLeftQuantity(0);
                order = orderRepository.save(order);
            }
            // Buy order has less stocks than needed
            else {
                // Create payment request before applying stocks transactions
                InsertPaymentRequest paymentRequest = paymentService.createBankTransfer(
                    order.getPrice().multiplyBy(buyOrder.getLeftQuantity()).toPaymentValue(),
                    buyDossier.getIban(),
                    dossier.getName() + ' ' + dossier.getSurname(),
                    dossier.getIban(),
                    String.format("Buy %d %s stocks for %s",
                        buyOrder.getLeftQuantity(),
                        order.getIsin(),
                        order.getPrice().toString()
                    )
                );
                ResponseEntity<Payment> res = paymentService.insertPayment(bank.externalApiToken, paymentRequest);
                if (res.getStatusCode() != HttpStatus.CREATED) {
                    System.out.printf("Payment request received status code %d\n", res.getStatusCode().value());
                    return;
                }
                System.out.printf("Payment request created successufully\n");

                // Update current orders
                order.setLeftQuantity(order.getLeftQuantity() - buyOrder.getLeftQuantity());
                order = orderRepository.save(order);

                buyOrder.setLeftQuantity(0);
                orderRepository.save(buyOrder);

                // Get best buy orders to start transactions with
                buyOrders = orderRepository.findBestBuyOrders(order.getIsin());
                if (buyOrders.isEmpty()) {
                    System.out.printf("No available buy orders with isin %s found\n", order.getIsin());
                    return;
                }
                buyOrder = buyOrders.getFirst();

                // Get dossier associated with buy order
                buyDossierOptional = dossierRepository.findById(buyOrder.getDossier());
                if (buyDossierOptional.isEmpty()) {
                    System.out.printf("No dossiers with id %s found\n", buyOrder.getDossier());
                    return;
                }
                buyDossier = buyDossierOptional.get();
            }
        }
    }
}
