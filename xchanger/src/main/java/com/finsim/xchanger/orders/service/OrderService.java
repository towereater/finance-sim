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

        return insertOrder(order);
    }

    public Order insertOrder(Order order) {
        return orderRepository.insert(order);
    }

    public void deleteOrder(String id) {
        orderRepository.deleteById(id);
    }

    public void elaborateOrder(Order order) {
        if (order.getType().equals("BUY")) {
            elaborateBuyOrder(order);
        } else if (order.getType().equals("SELL")) {
            elaborateSellOrder(order);
        }
    }

    public void elaborateBuyOrder(Order order) {
        List<Order> sellOrders = orderRepository.findBestSellOrders(order.getIsin());
        if (sellOrders.isEmpty()) {
            return;
        }
        Order sellOrder = sellOrders.getFirst();

        Optional<Dossier> dossierOptional = dossierRepository.findById(order.getDossier());
        Dossier dossier = dossierOptional.get();

        Optional<Bank> bankOptional = bankRepository.findByAbi(dossier.getAbi());
        Bank bank = bankOptional.get();

        Optional<Dossier> sellDossierOptional = dossierRepository.findById(sellOrder.getDossier());
        Dossier sellDossier = sellDossierOptional.get();

        while (order.getLeftQuantity() > 0
            && sellOrder != null
            && order.getPrice().getAmount() >= sellOrder.getPrice().getAmount()) {
            if (order.getLeftQuantity() <= sellOrder.getLeftQuantity()) {
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
                    return;
                }

                sellOrder.setLeftQuantity(sellOrder.getLeftQuantity() - order.getLeftQuantity());
                orderRepository.save(sellOrder);

                order.setLeftQuantity(0);
                order = orderRepository.save(order);
            } else {
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
                    return;
                }

                order.setLeftQuantity(order.getLeftQuantity() - sellOrder.getLeftQuantity());
                order = orderRepository.save(order);

                sellOrder.setLeftQuantity(0);
                orderRepository.save(sellOrder);

                sellOrders = orderRepository.findBestSellOrders(order.getIsin());
                if (sellOrders.isEmpty()) {
                    return;
                }
                sellOrder = sellOrders.getFirst();

                sellDossierOptional = dossierRepository.findById(sellOrder.getDossier());
                sellDossier = sellDossierOptional.get();
            }
        }
    }

    public void elaborateSellOrder(Order order) {
        List<Order> buyOrders = orderRepository.findBestBuyOrders(order.getIsin());
        if (buyOrders.isEmpty()) {
            return;
        }
        Order buyOrder = buyOrders.getFirst();

        Optional<Dossier> dossierOptional = dossierRepository.findById(order.getDossier());
        Dossier dossier = dossierOptional.get();

        Optional<Bank> bankOptional = bankRepository.findByAbi(dossier.getAbi());
        Bank bank = bankOptional.get();

        Optional<Dossier> buyDossierOptional = dossierRepository.findById(buyOrder.getDossier());
        Dossier buyDossier = buyDossierOptional.get();

        while (order.getLeftQuantity() > 0
            && buyOrder != null
            && order.getPrice().getAmount() <= buyOrder.getPrice().getAmount()) {
            if (order.getLeftQuantity() <= buyOrder.getLeftQuantity()) {
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
                    return;
                }

                buyOrder.setLeftQuantity(buyOrder.getLeftQuantity() - order.getLeftQuantity());
                orderRepository.save(buyOrder);

                order.setLeftQuantity(0);
                order = orderRepository.save(order);
            } else {
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
                    return;
                }

                order.setLeftQuantity(order.getLeftQuantity() - buyOrder.getLeftQuantity());
                order = orderRepository.save(order);

                buyOrder.setLeftQuantity(0);
                orderRepository.save(buyOrder);

                buyOrders = orderRepository.findBestBuyOrders(order.getIsin());
                if (buyOrders.isEmpty()) {
                    return;
                }
                buyOrder = buyOrders.getFirst();

                buyDossierOptional = dossierRepository.findById(buyOrder.getDossier());
                buyDossier = buyDossierOptional.get();
            }
        }
    }
}
