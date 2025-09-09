package com.finsim.xchanger.dossiers.model;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Document(collection = "dossiers")
public class Dossier {
    @Id
    private String id;

    private String name;
    private String surname;
    private String birth;

    public DossierDto toDto() {
        return new DossierDto(
            this.id,
            this.name,
            this.surname,
            this.birth
        );
    }
}
