package com.finsim.xchanger.banks.model;

import org.springframework.data.mongodb.core.index.Indexed;
import org.springframework.data.mongodb.core.mapping.Document;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Document(collection = "banks")
public class Bank {
    @Indexed(unique = true)
    private String abi;

    @Indexed(unique = true)
    private String apiToken;

    private String externalApiToken;

    public BankDto toDto() {
        return new BankDto(
            this.abi,
            this.apiToken,
            this.externalApiToken
        );
    }
}
