package com.finsim.xchanger.dossiers.model;

import java.util.List;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.index.Indexed;
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

    private String abi;
    private String externalId;
    @Indexed(unique = true)
    private String iban;

    private List<DossierStock> stocks;
    private DossierValue value;

    public DossierDto toDto() {
        return new DossierDto(
            this.id,
            this.name,
            this.surname,
            this.birth,
            this.abi,
            this.externalId,
            this.iban,
            this.stocks,
            this.value
        );
    }
}
