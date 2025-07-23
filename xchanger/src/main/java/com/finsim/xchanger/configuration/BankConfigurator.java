package com.finsim.xchanger.configuration;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

import com.finsim.xchanger.model.Bank;
import com.finsim.xchanger.service.BankService;

@Component
public class BankConfigurator implements CommandLineRunner {
	@Autowired
	private BankService bankService;
    
	@Override
	public void run(String... args) throws Exception {
		// Logging
		System.err.printf("Bank collection configuration started%n");

		if (bankService.count() == 0) {
			bankService.insertBank(new Bank("", "06270", "AAAAA"));
		}

		// Logging
		System.err.printf("Bank collection configuration ended%n");
	}
}
