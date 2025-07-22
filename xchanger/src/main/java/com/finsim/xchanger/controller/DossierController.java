package com.finsim.xchanger.controller;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import com.finsim.xchanger.entity.Dossier;
import com.finsim.xchanger.service.DossierService;

@RestController
@RequestMapping("/dossiers")
public class DossierController {
    @Autowired
    private DossierService dossierService;
    
    // Aggregate entity

    @GetMapping
    public List<Dossier> getAllDossiers() {
        return dossierService.findAllDossiers();
    }

    @PostMapping
    public ResponseEntity<Dossier> createDossier(@RequestBody Dossier dossier) {
        Dossier insertedDossier = dossierService.insertDossier(dossier);
        return new ResponseEntity<>(insertedDossier, HttpStatus.CREATED);
    }
    
    // Single entity

    @GetMapping("/{id}")
    public ResponseEntity<Dossier> getDossierById(@PathVariable String id) {
        return dossierService.findDossierById(id)
            .map(ResponseEntity::ok)
            .orElse(ResponseEntity.notFound().build());
    }
    
    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteDossier(@PathVariable String id) {
        dossierService.deleteDossier(id);
        return ResponseEntity.noContent().build();
    }
}