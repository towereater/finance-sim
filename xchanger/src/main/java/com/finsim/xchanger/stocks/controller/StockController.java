package com.finsim.xchanger.stocks.controller;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.finsim.xchanger.stocks.model.InsertStockRequest;
import com.finsim.xchanger.stocks.model.Stock;
import com.finsim.xchanger.stocks.model.StockDto;
import com.finsim.xchanger.stocks.service.StockService;

import jakarta.validation.Valid;

@RestController
@RequestMapping("/stocks")
public class StockController {
    @Autowired
    private StockService stockService;
    
    // Aggregate entity

    @GetMapping
    public ResponseEntity<List<StockDto>> getAllStocks(
        @RequestParam(defaultValue = "0") int page,
        @RequestParam(defaultValue = "50") int size
    ) {
        Pageable pageable = PageRequest.of(page, size);
        List<Stock> stocks = stockService.findAllStocks(pageable).getContent();

        if (stocks.isEmpty()) {
            return new ResponseEntity<>(HttpStatus.NOT_FOUND);
        } else {
            return new ResponseEntity<>(
                stocks.stream().map(stock -> stock.toDto()).toList(),
                HttpStatus.OK
            );
        }
    }

    @PostMapping
    public ResponseEntity<StockDto> createStock(
        @Valid @RequestBody InsertStockRequest request
    ) {
        Stock stock = stockService.createStock(request);
        return new ResponseEntity<>(stock.toDto(), HttpStatus.CREATED);
    }
    
    // Single entity

    @GetMapping("/{isin}")
    public ResponseEntity<StockDto> getStockByIsin(
        @PathVariable String isin
    ) {
        return stockService.findStockByIsin(isin)
            .map(stock -> new ResponseEntity<>(stock.toDto(), HttpStatus.OK))
            .orElse(ResponseEntity.notFound().build());
    }
}
