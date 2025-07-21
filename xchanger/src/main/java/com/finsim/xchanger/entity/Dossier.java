package com.finsim.xchanger.entity;

import org.springframework.data.annotation.Id;

public class Dossier {
    @Id
    public String id;

    public String name;
    public String surname;
    public String birth;

    public Dossier() {}

    public Dossier(String name, String surname, String birth) {
        this.name = name;
        this.surname = surname;
        this.birth = birth;
    }
}