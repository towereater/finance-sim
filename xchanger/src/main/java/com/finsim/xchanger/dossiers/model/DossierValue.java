package com.finsim.xchanger.dossiers.model;

import com.finsim.xchanger.common.model.Price;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class DossierValue {
    private Price value;
    private String timestamp;
}
