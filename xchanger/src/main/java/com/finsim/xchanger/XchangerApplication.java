package com.finsim.xchanger;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import com.finsim.xchanger.db.BankRepository;
import com.finsim.xchanger.entity.Bank;

@SpringBootApplication
public class XchangerApplication implements CommandLineRunner {
	@Autowired
	private BankRepository bankRepository;

	public static void main(String[] args) {
		SpringApplication.run(XchangerApplication.class, args);
	}

	@Override
	public void run(String... args) throws Exception {
		// Configure bank collection
		if (bankRepository.count() == 0) {
			bankRepository.insert(new Bank("06270", "AAAAA"));
		}
	}
}
