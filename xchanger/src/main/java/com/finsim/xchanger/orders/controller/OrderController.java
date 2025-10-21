package com.finsim.xchanger.orders.controller;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.finsim.xchanger.orders.model.InsertOrderRequest;
import com.finsim.xchanger.orders.model.Order;
import com.finsim.xchanger.orders.model.OrderDto;
import com.finsim.xchanger.orders.service.OrderService;

import jakarta.validation.Valid;

@RestController
@RequestMapping("/orders")
public class OrderController {
    @Autowired
    private OrderService orderService;
    
    // Aggregate entity

    @GetMapping
    public ResponseEntity<List<OrderDto>> getAllOrders(
        @RequestParam(defaultValue = "0") int page,
        @RequestParam(defaultValue = "50") int size
    ) {
        Pageable pageable = PageRequest.of(page, size);
        List<Order> orders = orderService.findAllOrders(pageable).getContent();

        if (orders.isEmpty()) {
            return new ResponseEntity<>(HttpStatus.NOT_FOUND);
        } else {
            return new ResponseEntity<>(
                orders.stream().map(order -> order.toDto()).toList(),
                HttpStatus.OK
            );
        }
    }

    @PostMapping
    public ResponseEntity<OrderDto> createOrder(
        @Valid @RequestBody InsertOrderRequest request
    ) {
        Order order = orderService.createOrder(request);
        return new ResponseEntity<>(order.toDto(), HttpStatus.CREATED);
    }
    
    // Single entity

    @GetMapping("/{id}")
    public ResponseEntity<OrderDto> getOrderById(
        @PathVariable String id
    ) {
        return orderService.findOrderById(id)
            .map(order -> new ResponseEntity<>(order.toDto(), HttpStatus.OK))
            .orElse(ResponseEntity.notFound().build());
    }
    
    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteOrder(
        @PathVariable String id
    ) {
        orderService.deleteOrder(id);
        return ResponseEntity.noContent().build();
    }
}
