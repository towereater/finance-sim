package com.finsim.xchanger.dossiers.model;

import java.util.List;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class DossierDto {
    private String id;

    private String name;
    private String surname;
    private String birth;

    private String abi;
    private String externalId;
    private String iban;
    
    private List<DossierStock> stocks;
    private DossierValue value;
}
