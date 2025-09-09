package com.finsim.xchanger.banks.service;

import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.finsim.xchanger.banks.model.Bank;
import com.finsim.xchanger.banks.repository.BankRepository;

@Service
public class BankService {
    @Autowired
    private BankRepository bankRepository;

    public Optional<Bank> findBankByApiToken(String apiToken) {
        return bankRepository.findByApiToken(apiToken);
    }

    public long count() {
        return bankRepository.count();
    }

    public Bank insertBank(Bank bank) {
        return bankRepository.insert(bank);
    }
}
