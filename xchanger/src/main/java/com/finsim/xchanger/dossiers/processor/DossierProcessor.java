package com.finsim.xchanger.dossiers.processor;

import java.time.Instant;
import java.util.List;
import java.util.Optional;
import java.util.concurrent.TimeUnit;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.PageRequest;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

import com.finsim.xchanger.common.model.Price;
import com.finsim.xchanger.dossiers.model.Dossier;
import com.finsim.xchanger.dossiers.model.DossierStock;
import com.finsim.xchanger.dossiers.model.DossierValue;
import com.finsim.xchanger.dossiers.service.DossierService;
import com.finsim.xchanger.stocks.model.Stock;
import com.finsim.xchanger.stocks.service.StockService;

@Component
public class DossierProcessor {
    @Autowired
    private DossierService dossierService;

    @Autowired
    private StockService stockService;
    
    @Scheduled(timeUnit = TimeUnit.MINUTES, initialDelay = 3, fixedDelay = 1)
    public void updateDossierValues() {
        System.out.printf("Update dossier job started\n");

        int pageNumber = 0;
        int pageSize = 10;

        List<Dossier> dossiers = dossierService.findAllDossiers(PageRequest.of(pageNumber, pageSize)).getContent();
        System.out.printf("Found %d dossiers\n", dossiers.size());
        while (dossiers.size() > 0) {
            for (Dossier dossier : dossiers) {
                List<DossierStock> dossierStocks = dossier.getStocks();
                if (dossierStocks == null) {
                    System.out.printf("Dossier %s has no stocks\n", dossier.getId());

                    dossier.setValue(new DossierValue(new Price(0, "EUR"), Instant.now().toString()));
                    dossierService.updateDossier(dossier);
                    continue;
                }

                float totalValue = 0;
                for (DossierStock dossierStock : dossierStocks) {
                    Optional<Stock> stockOptional = stockService.findStockByIsin(dossierStock.getIsin());
                    if (stockOptional.isEmpty()) {
                        System.out.printf("No stock with isin %s found\n", dossierStock.getIsin());
                        continue;
                    }

                    Stock stock = stockOptional.get();
                    if (stock.getPrices() != null && stock.getPrices().getDailyLast() != null) {
                        totalValue += stock.getPrices().getDailyLast().getAmount() * dossierStock.getTotal();
                    }
                }

                System.out.printf("Dossier %s has %f stocks total value\n", dossier.getId(), totalValue);
                dossier.setValue(new DossierValue(new Price(totalValue, "EUR"), Instant.now().toString()));
                dossierService.updateDossier(dossier);
            }

            pageNumber++;
            dossiers = dossierService.findAllDossiers(PageRequest.of(pageNumber, pageSize)).getContent();
        }

        System.out.printf("Update dossier job ended\n");
    }
}
