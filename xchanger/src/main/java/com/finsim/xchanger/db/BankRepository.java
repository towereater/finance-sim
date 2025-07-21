package com.finsim.xchanger.db;

import org.springframework.data.mongodb.repository.MongoRepository;

import com.finsim.xchanger.entity.Bank;

public interface BankRepository extends MongoRepository<Bank, String> {
  public Bank findByAbi(String abi);
  public Bank findByApiToken(String apiToken);
}
