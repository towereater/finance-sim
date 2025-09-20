package com.finsim.xchanger.dossiers.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

import com.finsim.xchanger.dossiers.model.Dossier;
import com.finsim.xchanger.dossiers.model.InsertDossierRequest;
import com.finsim.xchanger.dossiers.repository.DossierRepository;

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

        return insertDossier(dossier);
    }

    public Dossier insertDossier(Dossier dossier) {
        return dossierRepository.insert(dossier);
    }

    public void deleteDossier(String id) {
        dossierRepository.deleteById(id);
    }
}
