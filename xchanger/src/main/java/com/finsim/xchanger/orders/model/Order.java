package com.finsim.xchanger.orders.model;

import org.springframework.data.annotation.Id;

import com.finsim.xchanger.common.model.Price;

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

    public OrderDto toDto() {
        return new OrderDto(
            this.id,
            this.dossier,
            this.isin,
            this.type,
            this.price,
            this.quantity,
            this.options
        );
    }
}
