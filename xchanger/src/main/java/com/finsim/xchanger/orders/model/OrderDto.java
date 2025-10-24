package com.finsim.xchanger.orders.model;

import java.util.List;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.finsim.xchanger.common.model.Price;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class OrderDto {
    private String id;

    private String dossier;
    private String isin;
    private String type;

    private Price price;
    private int quantity;
    private String options;

    private int leftQuantity;
    @JsonInclude(Include.NON_NULL)
    private List<OrderTransaction> orderTransactions;
}
