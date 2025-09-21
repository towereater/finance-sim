package com.finsim.xchanger.payments.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class Payee {
    private String name;
    private AccountIdentification accountIdentification;
}
