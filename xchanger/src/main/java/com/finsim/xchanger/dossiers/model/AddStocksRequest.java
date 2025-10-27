package com.finsim.xchanger.dossiers.model;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;

public class AddStocksRequest {
    @NotBlank
    public String isin;
    @NotNull
    public int quantity;
}
