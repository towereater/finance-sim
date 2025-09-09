package com.finsim.xchanger.dossiers.model;

import jakarta.validation.constraints.NotBlank;

public class InsertDossierRequest {
    @NotBlank
    public String name;
    @NotBlank
    public String surname;
    @NotBlank
    public String birth;
}
