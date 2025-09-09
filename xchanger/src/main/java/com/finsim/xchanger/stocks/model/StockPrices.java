package com.finsim.xchanger.stocks.model;

import com.finsim.xchanger.common.model.Price;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class StockPrices {
    public Price dailyMax;
    public Price dailyMin;
    public Price dailyOpening;
}
