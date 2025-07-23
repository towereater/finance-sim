package com.finsim.xchanger.repository;

import java.util.Optional;

import org.springframework.data.mongodb.repository.MongoRepository;

import com.finsim.xchanger.model.Bank;

public interface BankRepository extends MongoRepository<Bank, String> {
  public Optional<Bank> findByAbi(String abi);
  public Optional<Bank> findByApiToken(String apiToken);
}
