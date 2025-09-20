package com.finsim.xchanger.banks.repository;

import java.util.Optional;

import org.springframework.data.mongodb.repository.MongoRepository;

import com.finsim.xchanger.banks.model.Bank;

public interface BankRepository extends MongoRepository<Bank, String> {
    public Optional<Bank> findByAbi(String abi);
    public Optional<Bank> findByApiToken(String apiToken);
    public void deleteByAbi(String abi);
}
