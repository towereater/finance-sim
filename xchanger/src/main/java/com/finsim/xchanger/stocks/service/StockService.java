package com.finsim.xchanger.stocks.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

import com.finsim.xchanger.stocks.model.InsertStockRequest;
import com.finsim.xchanger.stocks.model.Stock;
import com.finsim.xchanger.stocks.repository.StockRepository;

import java.util.Optional;

@Service
public class StockService {
    @Autowired
    private StockRepository stockRepository;

    public Page<Stock> findAllStocks(Pageable pageable) {
        return stockRepository.findAll(pageable);
    }

    public Optional<Stock> findStockByIsin(String isin) {
        return stockRepository.findById(isin);
    }

    public Stock createStock(InsertStockRequest stockRequest) {
        Stock stock = new Stock();
        stock.setIsin(stockRequest.isin);
        stock.setSymbol(stockRequest.symbol);
        stock.setDescription(stockRequest.description);
        stock.setType(stockRequest.type);

        return insertStock(stock);
    }

    public Stock insertStock(Stock stock) {
        return stockRepository.insert(stock);
    }

    public void deleteStock(String isin) {
        stockRepository.deleteById(isin);
    }
}
