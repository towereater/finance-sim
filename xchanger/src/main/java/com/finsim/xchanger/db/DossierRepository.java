package com.finsim.xchanger.db;

import org.springframework.data.mongodb.repository.MongoRepository;

import com.finsim.xchanger.entity.Dossier;

public interface DossierRepository extends MongoRepository<Dossier, String> {
  
}
