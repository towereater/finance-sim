package com.finsim.xchanger.payments.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class InsertPaymentRequest {
    private String type;
    private PaymentValue value;
    private Payer payer;
    private Payee payee;
    private String details;
}
