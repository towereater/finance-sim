package com.finsim.xchanger.payments.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class Payer {
    private AccountIdentification accountIdentification;
}
