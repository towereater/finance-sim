package com.finsim.xchanger.orders.repository;

import java.util.List;

import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.data.mongodb.repository.Query;

import com.finsim.xchanger.orders.model.Order;

public interface OrderRepository extends MongoRepository<Order, String> {
    @Query(value="{ 'isin': ?0 }", hint="order-buy-index")
    List<Order> findBestBuyOrders(String isin);

    @Query(value="{ 'isin': ?0 }", hint="order-sell-index")
    List<Order> findBestSellOrders(String isin);
}
