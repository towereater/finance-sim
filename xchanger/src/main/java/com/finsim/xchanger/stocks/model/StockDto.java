package com.finsim.xchanger.stocks.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.finsim.xchanger.orders.model.Order;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class StockDto {
    public String isin;

    public String symbol;
    public String description;
    public String type;

    @JsonInclude(Include.NON_NULL)
    public StockPrices prices;

    @JsonInclude(Include.NON_NULL)
    public Order[] sellOrders;
    @JsonInclude(Include.NON_NULL)
    public Order[] buyOrders;
}
