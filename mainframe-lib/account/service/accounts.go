package service

import (
	"encoding/json"
	"fmt"
	"mainframe-lib/account/model"
	ccom "mainframe-lib/common/config"
	scom "mainframe-lib/common/service"
	"net/http"
)

func GetAccounts(service ccom.Service, auth string, filter model.Account, from string, limit int) ([]model.Account, int, error) {
	// Construct the request
	url := "/accounts"

	url = fmt.Sprintf("%s?limit=%d", url, limit)
	if from != "" {
		url = fmt.Sprintf("%s&from=%s", url, from)
	}
	if filter.Owner != "" {
		url = fmt.Sprintf("%s&owner=%s", url, filter.Owner)
	}
	if filter.Id.Service != "" {
		url = fmt.Sprintf("%s&service=%s", url, filter.Id.Service)
	}

	// Execute the request
	res, err := scom.ExecuteHttpRequest(service, http.MethodGet, url, auth, "")
	if err != nil {
		return []model.Account{}, http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode == http.StatusNotFound {
		return []model.Account{}, res.StatusCode, nil
	}
	if res.StatusCode != http.StatusOK {
		return []model.Account{}, res.StatusCode, fmt.Errorf("get accounts returned status %d", res.StatusCode)
	}

	// Parse the response
	var accounts []model.Account
	err = json.NewDecoder(res.Body).Decode(&accounts)
	if err != nil {
		return []model.Account{}, http.StatusInternalServerError, err
	}

	return accounts, res.StatusCode, nil
}

func InsertAccount(service ccom.Service, auth string, payload model.InsertAccountInput) (int, error) {
	// Construct the request
	url := "/accounts"

	// Execute the request
	res, err := scom.ExecuteHttpRequest(service, http.MethodPost, url, auth, payload)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode != http.StatusCreated {
		return res.StatusCode, fmt.Errorf("insert account returned status %d", res.StatusCode)
	}

	return res.StatusCode, nil
}

func DeleteAccount(service ccom.Service, auth string, accountId model.AccountId) (int, error) {
	// Construct the request
	url := fmt.Sprintf("/accounts/services/%s/accounts/%s", accountId.Service, accountId.Account)

	// Execute the request
	res, err := scom.ExecuteHttpRequest(service, http.MethodDelete, url, auth, "")
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode != http.StatusNoContent {
		return res.StatusCode, fmt.Errorf("delete account returned status %d", res.StatusCode)
	}

	return res.StatusCode, nil
}
