package com.finsim.xchanger.dossiers.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class DossierStock {
    private String isin;
    private int total;
    private int available;    
}
