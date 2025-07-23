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
			bankService.insertBank(new Bank("06270", "3cae43527ddbc85be07f711577e1fe48e0c5c70d6511432f2b0349f737a09d63"));
		}

		// Logging
		System.err.printf("Bank collection configuration ended%n");
	}
}
