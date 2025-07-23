package com.finsim.xchanger.repository;

import org.springframework.data.mongodb.repository.MongoRepository;

import com.finsim.xchanger.model.Dossier;

public interface DossierRepository extends MongoRepository<Dossier, String> {
  
}
