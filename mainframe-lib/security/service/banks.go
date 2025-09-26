package service

import (
	"encoding/json"
	"fmt"
	com "mainframe-lib/common/service"
	"mainframe-lib/security/model"
	"net/http"
)

func GetBankByAbi(host string, timeout int, auth string, abi string) (model.Bank, int, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/banks/%s", host, abi)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodGet, url, timeout, auth, "")
	if err != nil {
		return model.Bank{}, http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode == http.StatusNotFound {
		return model.Bank{}, res.StatusCode, nil
	}
	if res.StatusCode != http.StatusOK {
		return model.Bank{}, res.StatusCode, fmt.Errorf("get user returned status %d", res.StatusCode)
	}

	// Parse the response
	var bank model.Bank
	err = json.NewDecoder(res.Body).Decode(&bank)
	if err != nil {
		return model.Bank{}, http.StatusInternalServerError, err
	}

	return bank, res.StatusCode, nil
}
