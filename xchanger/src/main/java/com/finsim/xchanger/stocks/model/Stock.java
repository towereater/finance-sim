package com.finsim.xchanger.stocks.model;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import com.finsim.xchanger.orders.model.Order;

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

    public StockPrices prices;

    public Order[] sellOrders;
    public Order[] buyOrders;

    public StockDto toDto() {
        return new StockDto(
            this.isin,
            this.symbol,
            this.description,
            this.type,
            this.prices,
            this.sellOrders,
            this.buyOrders
        );
    }
}
