package com.finsim.xchanger.orders.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

import com.finsim.xchanger.orders.configuration.OrderProducer;
import com.finsim.xchanger.orders.model.InsertOrderRequest;
import com.finsim.xchanger.orders.model.Order;
import com.finsim.xchanger.orders.repository.OrderRepository;

import java.util.Optional;

@Service
public class OrderService {
    @Autowired
    private OrderRepository orderRepository;

    @Autowired
    private OrderProducer orderProducer;

    public Page<Order> findAllOrders(Pageable pageable) {
        return orderRepository.findAll(pageable);
    }

    public Optional<Order> findOrderById(String id) {
        return orderRepository.findById(id);
    }

    public Optional<Order> findBestSellOrderWithPriceLowerThan(String isin, float amount) {
        Page<Order> orders = orderRepository.findBestSellOrdersWithPriceLowerThan(
            isin, amount, PageRequest.of(0, 1));
        
        return orders.isEmpty() ? Optional.empty() : Optional.of(orders.getContent().getFirst());
    }

    public Optional<Order> findBestBuyOrderWithPriceGreaterThan(String isin, float amount) {
        Page<Order> orders =  orderRepository.findBestBuyOrdersWithPriceGreaterThan(
            isin, amount, PageRequest.of(0, 1));
        
        return orders.isEmpty() ? Optional.empty() : Optional.of(orders.getContent().getFirst());
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

    public Order updateOrder(Order order) {
        return orderRepository.save(order);
    }

    public void deleteOrder(String id) {
        orderRepository.deleteById(id);
    }
}
