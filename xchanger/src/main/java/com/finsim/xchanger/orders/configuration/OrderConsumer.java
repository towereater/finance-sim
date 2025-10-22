package com.finsim.xchanger.orders.configuration;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

import com.finsim.xchanger.orders.service.OrderService;

@Component
public class OrderConsumer {
    @Autowired
    private OrderService orderService;

    @KafkaListener(topics = "${xchanger.kafka.topic.orders}")
    public void elaborateOrder(String id) {
        System.out.printf("Elaboration of order %s started\n", id);

        orderService.elaborateOrder(id);

        System.out.printf("Elaboration of order %s completed\n", id);
    }
}
