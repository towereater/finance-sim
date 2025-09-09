package com.finsim.xchanger.stocks.model;

import jakarta.validation.constraints.NotBlank;

public class InsertStockRequest {
    @NotBlank
    public String isin;

    @NotBlank
    public String symbol;
    @NotBlank
    public String description;
    @NotBlank
    public String type;
}
