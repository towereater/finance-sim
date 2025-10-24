package com.finsim.xchanger.orders.controller;

import java.util.List;
import java.util.Optional;

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

import com.finsim.xchanger.dossiers.model.Dossier;
import com.finsim.xchanger.dossiers.model.DossierStock;
import com.finsim.xchanger.dossiers.service.DossierService;
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

    @Autowired
    private DossierService dossierService;
    
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
        if (request.type.equals("SELL")) {
            // Get dossier associated with request
            Optional<Dossier> dossierOptional = dossierService.findDossierById(request.dossier);
            if (dossierOptional.isEmpty()) {
                System.out.printf("No dossiers with id %s found\n", request.dossier);
                return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
            }
            Dossier dossier = dossierOptional.get();

            if (dossier.getStocks() == null) {
                System.out.printf("Dossier with id %s has not available stocks\n", request.dossier);
                return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
            }

            // Check dossier stocks availability
            Optional<DossierStock> dossierStockOptional = dossier.getStocks().stream()
                .filter(ds -> ds.getIsin().equals(request.isin))
                .findFirst();
            if (dossierStockOptional.isEmpty()
            || dossierStockOptional.get().getAvailable() < request.quantity) {
                System.out.printf("Dossier with id %s has not enough stocks with isin %s\n",
                    request.dossier, request.isin);
                return new ResponseEntity<>(HttpStatus.BAD_REQUEST);
            }

            // Lock dossier stock availability
            DossierStock dossierStock = dossierStockOptional.get();
            dossierStock.setAvailable(dossierStock.getAvailable() - request.quantity);

            List<DossierStock> dossierStocks = dossier.getStocks();
            dossierStocks.replaceAll(ds -> ds.getIsin() == request.isin ? dossierStock : ds);

            dossier.setStocks(dossierStocks);
            dossierService.updateDossier(dossier);
        }

        // Create order and add it to queue for later processing
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
