package com.finsim.xchanger.controller;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.finsim.xchanger.model.Dossier;
import com.finsim.xchanger.model.InsertDossierRequest;
import com.finsim.xchanger.service.DossierService;

@RestController
@RequestMapping("/dossiers")
public class DossierController {
    @Autowired
    private DossierService dossierService;
    
    // Aggregate entity

    @GetMapping
    public ResponseEntity<List<Dossier>> getAllDossiers() {
        List<Dossier> dossiers = dossierService.findAllDossiers();

        if (dossiers.isEmpty()) {
            return new ResponseEntity<>(HttpStatus.NOT_FOUND);
        } else {
            return new ResponseEntity<>(dossiers, HttpStatus.OK);
        }
    }

    @PostMapping
    public ResponseEntity<Dossier> createDossier(@RequestBody InsertDossierRequest request) {
        Dossier insertedDossier = dossierService.insertDossier(
            new Dossier(request.name, request.surname, request.birth));
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
