package com.finsim.xchanger.banks.model;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Size;

public class InsertBankRequest {
    @Size(min = 5, max = 5)
    public String abi;
    @NotBlank
    public String apiToken;
    public String externalApiToken;
}
