package com.finsim.xchanger.orders.repository;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.data.mongodb.repository.Query;

import com.finsim.xchanger.orders.model.Order;

public interface OrderRepository extends MongoRepository<Order, String> {
    @Query(value="{ 'isin': ?0 }", hint="order-buy-index")
    Page<Order> findBestBuyOrders(String isin, Pageable pageable);

    @Query(value="{ 'isin': ?0, 'price.amount': { '$gte': ?1 } }", hint="order-buy-index")
    Page<Order> findBestBuyOrdersWithPriceGreaterThan(String isin, float amount, Pageable pageable);

    @Query(value="{ 'isin': ?0 }", hint="order-sell-index")
    Page<Order> findBestSellOrders(String isin, Pageable pageable);

    @Query(value="{ 'isin': ?0, 'price.amount': { '$lte': ?1 } }", hint="order-sell-index")
    Page<Order> findBestSellOrdersWithPriceLowerThan(String isin, float amount, Pageable pageable);
}
