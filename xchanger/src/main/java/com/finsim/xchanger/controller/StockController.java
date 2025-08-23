package com.finsim.xchanger.controller;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.finsim.xchanger.model.Stock;
import com.finsim.xchanger.service.StockService;

@RestController
@RequestMapping("/stocks")
public class StockController {
    @Autowired
    private StockService stockService;
    
    // Aggregate entity

    @GetMapping
    public ResponseEntity<List<Stock>> getAllStocks() {
        List<Stock> stocks = stockService.findAllStocks();

        if (stocks.isEmpty()) {
            return new ResponseEntity<>(HttpStatus.NOT_FOUND);
        } else {
            return new ResponseEntity<>(stocks, HttpStatus.OK);
        }
    }

    @PostMapping
    public ResponseEntity<Stock> createStock(@RequestBody Stock stock) {
        Stock insertedDossier = stockService.insertStock(stock);
        return new ResponseEntity<>(insertedDossier, HttpStatus.CREATED);
    }
    
    // Single entity

    @GetMapping("/{isin}")
    public ResponseEntity<Stock> getStockByIsin(@PathVariable String isin) {
        return stockService.findStockByIsin(isin)
            .map(ResponseEntity::ok)
            .orElse(ResponseEntity.notFound().build());
    }
}
