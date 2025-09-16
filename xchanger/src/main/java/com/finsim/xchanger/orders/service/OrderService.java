package com.finsim.xchanger.orders.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

import com.finsim.xchanger.orders.model.InsertOrderRequest;
import com.finsim.xchanger.orders.model.Order;
import com.finsim.xchanger.orders.repository.OrderRepository;

import java.util.List;
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

        while (order.getLeftQuantity() > 0
            && sellOrder != null
            && order.getPrice().getAmount() >= sellOrder.getPrice().getAmount()) {
            if (order.getLeftQuantity() <= sellOrder.getLeftQuantity()) {
                sellOrder.setLeftQuantity(sellOrder.getLeftQuantity() - order.getLeftQuantity());
                orderRepository.save(sellOrder);

                order.setLeftQuantity(0);
                order = orderRepository.save(order);
            } else {
                order.setLeftQuantity(order.getLeftQuantity() - sellOrder.getLeftQuantity());
                order = orderRepository.save(order);

                sellOrder.setLeftQuantity(0);
                orderRepository.save(sellOrder);

                sellOrders = orderRepository.findBestSellOrders(order.getIsin());
                if (sellOrders.isEmpty()) {
                    return;
                }
                sellOrder = sellOrders.getFirst();
            }
        }
    }

    public void elaborateSellOrder(Order order) {
        List<Order> buyOrders = orderRepository.findBestBuyOrders(order.getIsin());
        if (buyOrders.isEmpty()) {
            return;
        }
        Order buyOrder = buyOrders.getFirst();

        while (order.getLeftQuantity() > 0
            && buyOrder != null
            && order.getPrice().getAmount() <= buyOrder.getPrice().getAmount()) {
            if (order.getLeftQuantity() <= buyOrder.getLeftQuantity()) {
                buyOrder.setLeftQuantity(buyOrder.getLeftQuantity() - order.getLeftQuantity());
                orderRepository.save(buyOrder);

                order.setLeftQuantity(0);
                order = orderRepository.save(order);
            } else {
                order.setLeftQuantity(order.getLeftQuantity() - buyOrder.getLeftQuantity());
                order = orderRepository.save(order);

                buyOrder.setLeftQuantity(0);
                orderRepository.save(buyOrder);

                buyOrders = orderRepository.findBestBuyOrders(order.getIsin());
                if (buyOrders.isEmpty()) {
                    return;
                }
                buyOrder = buyOrders.getFirst();
            }
        }
    }
}
