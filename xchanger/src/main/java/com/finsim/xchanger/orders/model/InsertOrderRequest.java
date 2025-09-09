package com.finsim.xchanger.orders.model;

import com.finsim.xchanger.common.model.Price;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;

public class InsertOrderRequest {
    @NotBlank
    public String dossier;
    @NotBlank
    public String isin;
    @NotBlank
    public String type;

    @NotNull
    public Price price;
    @NotNull
    public int quantity;
    @NotBlank
    public String options;
}
