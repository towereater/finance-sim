package com.finsim.xchanger.banks.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.finsim.xchanger.banks.service.BankService;
import com.finsim.xchanger.banks.model.Bank;
import com.finsim.xchanger.banks.model.BankDto;
import com.finsim.xchanger.banks.model.InsertBankRequest;

import jakarta.validation.Valid;

@RestController
@RequestMapping("/banks")
public class BankController {
    @Autowired
    private BankService bankService;
    
    // Aggregate entity

    @PostMapping
    public ResponseEntity<BankDto> createBank(
        @Valid @RequestBody InsertBankRequest request
    ) {
        Bank bank = bankService.createBank(request);
        return new ResponseEntity<>(bank.toDto(), HttpStatus.CREATED);
    }
    
    // Single entity

    @GetMapping("/{abi}")
    public ResponseEntity<BankDto> getBankByAbi(
        @PathVariable String abi
    ) {
        return bankService.findBankByAbi(abi)
            .map(bank -> new ResponseEntity<>(bank.toDto(), HttpStatus.OK))
            .orElse(ResponseEntity.notFound().build());
    }
    
    @DeleteMapping("/{abi}")
    public ResponseEntity<Void> deleteBank(
        @PathVariable String abi
    ) {
        bankService.deleteBank(abi);
        return ResponseEntity.noContent().build();
    }
}
