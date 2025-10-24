package com.finsim.xchanger.stocks.model;

import com.finsim.xchanger.common.model.Price;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class StockPrices {
    private Price dailyMax;
    private Price dailyMin;
    private Price dailyOpening;
    private Price dailyLast;
}
