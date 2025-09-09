package com.finsim.xchanger.stocks.repository;

import org.springframework.data.mongodb.repository.MongoRepository;

import com.finsim.xchanger.stocks.model.Stock;

public interface StockRepository extends MongoRepository<Stock, String> {
}
