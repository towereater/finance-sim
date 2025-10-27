package com.finsim.xchanger.dossiers.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

import com.finsim.xchanger.common.model.Price;
import com.finsim.xchanger.dossiers.model.Dossier;
import com.finsim.xchanger.dossiers.model.DossierStock;
import com.finsim.xchanger.dossiers.model.DossierValue;
import com.finsim.xchanger.dossiers.model.InsertDossierRequest;
import com.finsim.xchanger.dossiers.repository.DossierRepository;

import java.time.Instant;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;
import java.util.Optional;

@Service
public class DossierService {
    @Autowired
    private DossierRepository dossierRepository;

    public Page<Dossier> findAllDossiers(Pageable pageable) {
        return dossierRepository.findAll(pageable);
    }

    public Optional<Dossier> findDossierById(String id) {
        return dossierRepository.findById(id);
    }

    public Dossier createDossier(InsertDossierRequest dossierRequest, String abi) {
        Dossier dossier = new Dossier();
        dossier.setName(dossierRequest.name);
        dossier.setSurname(dossierRequest.surname);
        dossier.setBirth(dossierRequest.birth);
        dossier.setAbi(abi);
        dossier.setExternalId(dossierRequest.externalId);
        dossier.setIban(dossierRequest.iban);
        dossier.setValue(new DossierValue(new Price(0, "EUR"), Instant.now().toString()));

        return insertDossier(dossier);
    }

    public Dossier insertDossier(Dossier dossier) {
        return dossierRepository.insert(dossier);
    }

    public Dossier updateDossier(Dossier dossier) {
        return dossierRepository.save(dossier);
    }

    public Dossier addStocks(Dossier dossier, String isin, int quantity) {
        List<DossierStock> stocks = dossier.getStocks();

        if (stocks == null) {
            stocks = new ArrayList<DossierStock>();
            stocks.add(new DossierStock(isin, quantity, quantity));
        } else {
            Optional<DossierStock> stocksOptional = stocks.stream()
                .filter(ds -> ds.getIsin().equals(isin))
                .findFirst();

            if (stocksOptional.isEmpty()) {
                stocks.add(new DossierStock(isin, quantity, quantity));
                stocks.sort(new Comparator<DossierStock>() {
                    @Override
                    public int compare(DossierStock rhs, DossierStock lhs) {
                        return rhs.getIsin().compareTo(lhs.getIsin());
                    }
                });
            } else {
                DossierStock stock = stocksOptional.get();
                stock.setTotal(stock.getTotal() + quantity);
                stock.setAvailable(stock.getAvailable() + quantity);
                stocks.replaceAll(ds -> ds.getIsin().equals(isin) ? stock : ds);
            }
        }

        dossier.setStocks(stocks);
        return updateDossier(dossier);
    }

    public void deleteDossier(String id) {
        dossierRepository.deleteById(id);
    }
}
