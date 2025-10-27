package com.finsim.xchanger.stocks.repository;

import java.util.Optional;

import org.springframework.data.mongodb.repository.MongoRepository;

import com.finsim.xchanger.stocks.model.Stock;

public interface StockRepository extends MongoRepository<Stock, String> {
    public Optional<Stock> findByIsin(String isin);
    public void deleteByIsin(String isin);
}
