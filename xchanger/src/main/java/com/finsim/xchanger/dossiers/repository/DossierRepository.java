package com.finsim.xchanger.dossiers.repository;

import org.springframework.data.mongodb.repository.MongoRepository;

import com.finsim.xchanger.dossiers.model.Dossier;

public interface DossierRepository extends MongoRepository<Dossier, String> {
    
}
