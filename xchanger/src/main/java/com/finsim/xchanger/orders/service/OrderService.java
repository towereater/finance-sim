package com.finsim.xchanger.orders.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

import com.finsim.xchanger.orders.model.InsertOrderRequest;
import com.finsim.xchanger.orders.model.Order;
import com.finsim.xchanger.orders.repository.OrderRepository;

import java.util.Optional;

@Service
public class OrderService {
    @Autowired
    private OrderRepository orderRepository;

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
        order.setQuantity(order.quantity);
        order.setOptions(orderRequest.options);

        return order;
    }

    public Order insertOrder(Order order) {
        return orderRepository.insert(order);
    }

    public void deleteOrder(String id) {
        orderRepository.deleteById(id);
    }
}
