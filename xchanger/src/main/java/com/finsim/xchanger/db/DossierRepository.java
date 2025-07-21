package com.finsim.xchanger.db;

import java.util.List;

import org.springframework.data.mongodb.repository.MongoRepository;

import com.finsim.xchanger.entity.Dossier;

public interface DossierRepository extends MongoRepository<Dossier, String> {
  public List<Dossier> findByOwner(String owner);
}
