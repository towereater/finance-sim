package com.finsim.xchanger.model;

import org.springframework.data.annotation.Id;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class Order {
    @Id
    public String id;

    public String dossier;
    public String isin;
    public String type;

    public Price price;
    public int quantity;
    public String options;

    public Order(String dossier, String isin, String type, Price price, int quantity, String options) {
        this.dossier = dossier;
        this.isin = isin;
        this.type = type;

        this.price = price;
        this.quantity = quantity;
        this.options = options;
    }
}
