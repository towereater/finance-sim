package com.finsim.xchanger.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.finsim.xchanger.model.Dossier;
import com.finsim.xchanger.repository.DossierRepository;

import java.util.List;
import java.util.Optional;

@Service
public class DossierService {
    @Autowired
    private DossierRepository dossierRepository;

    public List<Dossier> findAllDossiers() {
        return dossierRepository.findAll();
    }

    public Optional<Dossier> findDossierById(String id) {
        return dossierRepository.findById(id);
    }

    public Dossier insertDossier(Dossier dossier) {
        return dossierRepository.insert(dossier);
    }

    public void deleteDossier(String id) {
        dossierRepository.deleteById(id);
    }
}