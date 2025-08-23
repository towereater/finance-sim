package com.finsim.xchanger.repository;

import org.springframework.data.mongodb.repository.MongoRepository;

import com.finsim.xchanger.model.Stock;

public interface StockRepository extends MongoRepository<Stock, String> {
  
}
