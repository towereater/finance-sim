package com.finsim.xchanger.entity;

import org.springframework.data.annotation.Id;

public class Bank {
    @Id
    public String id;

    public String abi;
    public String apiToken;

    public void setAbi(String abi) {
        this.abi = abi;
    }

    public String getAbi() {
        return abi;
    }

    public void setApiToken(String apiToken) {
        this.apiToken = apiToken;
    }

    public String getApiToken() {
        return apiToken;
    }

    public Bank() {}

    public Bank(String abi, String apiToken) {
        this.abi = abi;
        this.apiToken = apiToken;
    }
}
