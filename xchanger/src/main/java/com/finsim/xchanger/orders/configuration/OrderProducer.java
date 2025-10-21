package com.finsim.xchanger.orders.configuration;

import java.time.Instant;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

@Component
public class OrderProducer {
    @Value("${xchanger.kafka.topic.orders}")
    private String ordersTopic;

    @Autowired
    private KafkaTemplate<String, String> kafkaTemplate;

    public void queueOrder(String id) {
        kafkaTemplate.send(ordersTopic, Instant.now().toString(), id);
    }
}