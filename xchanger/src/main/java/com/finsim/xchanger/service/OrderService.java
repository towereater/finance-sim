package com.finsim.xchanger.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.finsim.xchanger.model.Order;
import com.finsim.xchanger.repository.OrderRepository;

import java.util.List;
import java.util.Optional;

@Service
public class OrderService {
    @Autowired
    private OrderRepository orderRepository;

    public List<Order> findAllOrders() {
        return orderRepository.findAll();
    }

    public Optional<Order> findOrderById(String id) {
        return orderRepository.findById(id);
    }

    public Order insertOrder(Order order) {
        return orderRepository.insert(order);
    }

    public void deleteOrder(String id) {
        orderRepository.deleteById(id);
    }
}
