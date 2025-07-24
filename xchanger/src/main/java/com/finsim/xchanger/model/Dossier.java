package com.finsim.xchanger.model;

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
    public String id;

    public String name;
    public String surname;
    public String birth;

    public Dossier(String name, String surname, String birth) {
        this.name = name;
        this.surname = surname;
        this.birth = birth;
    }
}
