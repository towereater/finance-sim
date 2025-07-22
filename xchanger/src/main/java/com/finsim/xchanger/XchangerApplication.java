package com.finsim.xchanger;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import com.finsim.xchanger.entity.Bank;
import com.finsim.xchanger.service.BankService;

@SpringBootApplication
public class XchangerApplication implements CommandLineRunner {
	@Autowired
	private BankService bankService;

	public static void main(String[] args) {
		SpringApplication.run(XchangerApplication.class, args);
	}

	@Override
	public void run(String... args) throws Exception {
		// Configure bank collection
		if (bankService.count() == 0) {
			bankService.insertBank(new Bank("06270", "AAAAA"));
		}
	}
}
