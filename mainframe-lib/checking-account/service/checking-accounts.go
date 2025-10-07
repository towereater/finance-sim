package service

import (
	"encoding/json"
	"fmt"
	"mainframe-lib/checking-account/model"
	ccom "mainframe-lib/common/config"
	scom "mainframe-lib/common/service"
	"net/http"
)

func GetAccount(service ccom.ServiceConfig, auth string, accountId string) (model.CheckingAccount, int, error) {
	// Construct the request
	url := fmt.Sprintf("/checking-accounts/%s", accountId)

	// Execute the request
	res, err := scom.ExecuteHttpRequest(service, http.MethodGet, url, auth, "")
	if err != nil {
		return model.CheckingAccount{}, http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode == http.StatusNotFound {
		return model.CheckingAccount{}, res.StatusCode, nil
	}
	if res.StatusCode != http.StatusOK {
		return model.CheckingAccount{}, res.StatusCode, fmt.Errorf("get checking account returned status %d", res.StatusCode)
	}

	// Parse the response
	var ckAccount model.CheckingAccount
	err = json.NewDecoder(res.Body).Decode(&ckAccount)
	if err != nil {
		return model.CheckingAccount{}, http.StatusInternalServerError, err
	}

	return ckAccount, res.StatusCode, nil
}
