package com.finsim.xchanger.orders.model;

import java.util.List;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import com.finsim.xchanger.common.model.Price;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Document(collection = "orders")
public class Order {
    @Id
    private String id;

    private String dossier;
    private String isin;
    private String type;

    private Price price;
    private int quantity;
    private String options;

    private int leftQuantity;
    private List<OrderTransaction> orderTransactions;

    public OrderDto toDto() {
        return new OrderDto(
            this.id,
            this.dossier,
            this.isin,
            this.type,
            this.price,
            this.quantity,
            this.options,
            this.leftQuantity,
            this.orderTransactions
        );
    }
}
