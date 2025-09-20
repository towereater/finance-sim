package com.finsim.xchanger.dossiers.controller;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.PageRequest;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestAttribute;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.finsim.xchanger.dossiers.model.Dossier;
import com.finsim.xchanger.dossiers.model.DossierDto;
import com.finsim.xchanger.dossiers.model.InsertDossierRequest;
import com.finsim.xchanger.dossiers.service.DossierService;

import jakarta.validation.Valid;

@RestController
@RequestMapping("/dossiers")
public class DossierController {
    @Autowired
    private DossierService dossierService;
    
    // Aggregate entity

    @GetMapping
    public ResponseEntity<List<DossierDto>> getAllDossiers(
        @RequestParam(defaultValue = "0") int page,
        @RequestParam(defaultValue = "50") int size
    ) {
        Pageable pageable = PageRequest.of(page, size);
        List<Dossier> dossiers = dossierService.findAllDossiers(pageable).getContent();

        if (dossiers.isEmpty()) {
            return new ResponseEntity<>(HttpStatus.NOT_FOUND);
        } else {
            return new ResponseEntity<>(
                dossiers.stream().map(dossier -> dossier.toDto()).toList(),
                HttpStatus.OK
            );
        }
    }

    @PostMapping
    public ResponseEntity<DossierDto> createDossier(
        @Valid @RequestBody InsertDossierRequest request,
        @RequestAttribute String abi
    ) {
        Dossier dossier = dossierService.createDossier(request, abi);
        return new ResponseEntity<>(dossier.toDto(), HttpStatus.CREATED);
    }
    
    // Single entity

    @GetMapping("/{id}")
    public ResponseEntity<DossierDto> getDossierById(
        @PathVariable String id
    ) {
        return dossierService.findDossierById(id)
            .map(dossier -> new ResponseEntity<>(dossier.toDto(), HttpStatus.OK))
            .orElse(ResponseEntity.notFound().build());
    }
    
    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteDossier(
        @PathVariable String id
    ) {
        dossierService.deleteDossier(id);
        return ResponseEntity.noContent().build();
    }
}
