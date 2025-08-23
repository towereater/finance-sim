package com.finsim.xchanger.model;

public class InsertOrderRequest {
    public String dossier;
    public String isin;
    public String type;

    public Price price;
    public int quantity;
    public String options;
}
