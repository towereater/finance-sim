package service

import (
	"encoding/json"
	"fmt"
	"mainframe-lib/checking-account/model"

	com "mainframe-lib/common/service"
	"net/http"
)

func GetAccount(host string, timeout int, auth string, accountId string) (model.CheckingAccount, int, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/checking-accounts/%s", host, accountId)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodGet, url, timeout, auth, "")
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
