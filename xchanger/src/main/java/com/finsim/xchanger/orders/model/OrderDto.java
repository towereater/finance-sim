package com.finsim.xchanger.orders.model;

import com.finsim.xchanger.common.model.Price;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class OrderDto {
    public String id;

    public String dossier;
    public String isin;
    public String type;

    public Price price;
    public int quantity;
    public String options;
}
