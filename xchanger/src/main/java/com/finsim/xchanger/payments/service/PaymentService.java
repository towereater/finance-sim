package com.finsim.xchanger.payments.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import com.finsim.xchanger.payments.model.AccountIdentification;
import com.finsim.xchanger.payments.model.InsertPaymentRequest;
import com.finsim.xchanger.payments.model.Payee;
import com.finsim.xchanger.payments.model.Payer;
import com.finsim.xchanger.payments.model.Payment;
import com.finsim.xchanger.payments.model.PaymentValue;

@Service
public class PaymentService {
    @Autowired
    private RestTemplate restTemplate;

    @Value("${xchanger.payments.host}")
    private String baseUrl;

    public InsertPaymentRequest createBankTransfer(PaymentValue value, String payerIban, String payeeName, String payeeIban, String details) {
        InsertPaymentRequest req = new InsertPaymentRequest();
        req.setType("BANK_TRANSFER");
        req.setValue(value);
        req.setPayer(new Payer(new AccountIdentification("IBAN", payerIban)));
        req.setPayee(new Payee(payeeName, new AccountIdentification("IBAN", payeeIban)));
        req.setDetails(details);

        return req;
    }

    public ResponseEntity<Payment> insertPayment(String auth, InsertPaymentRequest insertPaymentRequest) {
        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.APPLICATION_JSON);
        headers.set("Authorization", auth);

        HttpEntity<InsertPaymentRequest> req = new HttpEntity<>(insertPaymentRequest, headers);
        String url = baseUrl + "/payments";

        ResponseEntity<Payment> res = restTemplate.postForEntity(url, req, Payment.class);

        return res;
    }
}
