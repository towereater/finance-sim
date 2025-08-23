package com.finsim.xchanger.model;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import com.mongodb.lang.Nullable;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Document(collection = "stocks")
public class Stock {
    @Id
    public String isin;

    public String symbol;
    public String description;
    public String type;

    @Nullable
    public StockPrices prices;

    @Nullable
    public Order[] sellOrders;
    @Nullable
    public Order[] buyOrders;
}
