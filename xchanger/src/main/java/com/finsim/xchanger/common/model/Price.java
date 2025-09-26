package com.finsim.xchanger.common.model;

import com.finsim.xchanger.payments.model.PaymentValue;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class Price {
    public float amount;
    public String currency;

    public Price multiplyBy(float multiplier) {
        return new Price(this.amount * multiplier, this.currency);
    }

    public PaymentValue toPaymentValue() {
        return new PaymentValue(this.amount, this.currency);
    }

    @Override
    public String toString() {
        return String.format("%.2f %s", this.amount, this.currency);
    }
}
