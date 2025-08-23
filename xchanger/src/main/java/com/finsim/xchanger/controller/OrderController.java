package com.finsim.xchanger.controller;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.finsim.xchanger.model.Order;
import com.finsim.xchanger.model.InsertOrderRequest;
import com.finsim.xchanger.service.OrderService;

@RestController
@RequestMapping("/orders")
public class OrderController {
    @Autowired
    private OrderService orderService;
    
    // Aggregate entity

    @GetMapping
    public ResponseEntity<List<Order>> getAllOrders() {
        List<Order> dossiers = orderService.findAllOrders();

        if (dossiers.isEmpty()) {
            return new ResponseEntity<>(HttpStatus.NOT_FOUND);
        } else {
            return new ResponseEntity<>(dossiers, HttpStatus.OK);
        }
    }

    @PostMapping
    public ResponseEntity<Order> createOrder(@RequestBody InsertOrderRequest request) {
        Order insertedDossier = orderService.insertOrder(
            new Order(request.dossier, request.isin, request.type, request.price, request.quantity, request.options));
        return new ResponseEntity<>(insertedDossier, HttpStatus.CREATED);
    }
    
    // Single entity

    @GetMapping("/{id}")
    public ResponseEntity<Order> getOrderById(@PathVariable String id) {
        return orderService.findOrderById(id)
            .map(ResponseEntity::ok)
            .orElse(ResponseEntity.notFound().build());
    }
    
    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteOrder(@PathVariable String id) {
        orderService.deleteOrder(id);
        return ResponseEntity.noContent().build();
    }
}
