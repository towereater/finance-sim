package com.finsim.xchanger.stocks.model;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonInclude.Include;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class StockDto {
    private String isin;

    private String symbol;
    private String description;
    private String type;

    @JsonInclude(Include.NON_NULL)
    private StockPrices prices;
}
