package com.finsim.xchanger.model;

import org.springframework.data.mongodb.core.index.Indexed;
import org.springframework.data.mongodb.core.mapping.Document;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Document(collection = "banks")
public class Bank {
    @Indexed(unique = true)
    public String abi;

    @Indexed(unique = true)
    public String apiToken;
}
