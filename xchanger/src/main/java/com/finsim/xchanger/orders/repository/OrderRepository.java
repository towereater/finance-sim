package com.finsim.xchanger.orders.repository;

import org.springframework.data.mongodb.repository.MongoRepository;

import com.finsim.xchanger.orders.model.Order;

public interface OrderRepository extends MongoRepository<Order, String> {
}
