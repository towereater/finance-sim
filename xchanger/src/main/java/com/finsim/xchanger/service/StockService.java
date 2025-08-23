package com.finsim.xchanger.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.finsim.xchanger.model.Stock;
import com.finsim.xchanger.repository.StockRepository;

import java.util.List;
import java.util.Optional;

@Service
public class StockService {
    @Autowired
    private StockRepository stockRepository;

    public List<Stock> findAllStocks() {
        return stockRepository.findAll();
    }

    public Optional<Stock> findStockByIsin(String isin) {
        return stockRepository.findById(isin);
    }

    public Stock insertStock(Stock stock) {
        return stockRepository.insert(stock);
    }

    public void deleteStock(String isin) {
        stockRepository.deleteById(isin);
    }
}
