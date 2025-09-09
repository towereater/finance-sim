package com.finsim.xchanger.dossiers.model;

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
}
