package com.finsim.xchanger.orders.model;

import com.finsim.xchanger.common.model.Price;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class OrderTransaction {
    private String dossier;
    private int quantity;
    private Price price;
    private String timestamp;
}
