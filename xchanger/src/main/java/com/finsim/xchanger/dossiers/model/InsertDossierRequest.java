package com.finsim.xchanger.dossiers.model;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Size;

public class InsertDossierRequest {
    @NotBlank
    public String name;
    @NotBlank
    public String surname;
    @NotBlank
    public String birth;

    @NotBlank
    public String externalId;
    @Size(min = 27, max = 33)
    public String iban;
}
