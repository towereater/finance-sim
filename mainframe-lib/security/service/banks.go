package service

import (
	"encoding/json"
	"fmt"
	com "mainframe-lib/common/service"
	"mainframe-lib/security/model"
	"net/http"
)

func GetBankByAbi(host string, timeout int, auth string, abi string) (model.Bank, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/banks/%s", host, abi)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodGet, url, timeout, auth, "")
	if err != nil {
		return model.Bank{}, err
	}

	// Check response
	if res.StatusCode != http.StatusOK {
		return model.Bank{}, fmt.Errorf("get user returned status %d", res.StatusCode)
	}

	// Parse the response
	var bank model.Bank
	err = json.NewDecoder(res.Body).Decode(&bank)
	if err != nil {
		return model.Bank{}, err
	}

	return bank, nil
}
