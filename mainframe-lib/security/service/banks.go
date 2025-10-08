package service

import (
	"encoding/json"
	"fmt"
	ccom "mainframe-lib/common/config"
	scom "mainframe-lib/common/service"
	"mainframe-lib/security/model"
	"net/http"
)

func GetBankByAbi(service ccom.Service, auth string, abi string) (model.Bank, int, error) {
	// Construct the request
	url := fmt.Sprintf("/banks/%s", abi)

	// Execute the request
	res, err := scom.ExecuteHttpRequest(service, http.MethodGet, url, auth, "")
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
