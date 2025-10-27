package com.finsim.xchanger.stocks.model;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.index.Indexed;
import org.springframework.data.mongodb.core.mapping.Document;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Document(collection = "stocks")
public class Stock {
    @Id
    private String id;

    @Indexed(unique = true)
    private String isin;

    private String symbol;
    private String description;
    private String type;

    private StockPrices prices;

    public StockDto toDto() {
        return new StockDto(
            this.isin,
            this.symbol,
            this.description,
            this.type,
            this.prices
        );
    }
}
