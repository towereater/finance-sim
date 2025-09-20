package com.finsim.xchanger.banks.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class BankDto {
    private String abi;
    private String apiToken;
    private String externalApiToken;
}
