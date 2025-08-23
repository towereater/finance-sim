package com.finsim.xchanger.repository;

import org.springframework.data.mongodb.repository.MongoRepository;

import com.finsim.xchanger.model.Order;

public interface OrderRepository extends MongoRepository<Order, String> {
  
}
