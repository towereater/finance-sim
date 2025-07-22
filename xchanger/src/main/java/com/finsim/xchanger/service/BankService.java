package com.finsim.xchanger.service;

import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.finsim.xchanger.db.BankRepository;
import com.finsim.xchanger.entity.Bank;

@Service
public class BankService {
    @Autowired
    private BankRepository bankRepository;

    public Optional<Bank> findBankByApiToken(String apiToken) {
        return bankRepository.findByApiToken(apiToken);
    }
}
