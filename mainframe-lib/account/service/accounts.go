package service

import (
	"encoding/json"
	"fmt"
	"mainframe-lib/account/model"
	com "mainframe-lib/common/service"
	"net/http"
)

func GetAccounts(host string, timeout int, auth string, filter model.Account, from string, limit int) ([]model.Account, int, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/accounts", host)

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
	res, err := com.ExecuteHttpRequest(http.MethodGet, url, timeout, auth, "")
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

func InsertAccount(host string, timeout int, auth string, payload model.InsertAccountInput) (int, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/accounts", host)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodPost, url, timeout, auth, payload)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode != http.StatusCreated {
		return res.StatusCode, fmt.Errorf("insert account returned status %d", res.StatusCode)
	}

	return res.StatusCode, nil
}

func DeleteAccount(host string, timeout int, auth string, accountId model.AccountId) (int, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/accounts/services/%s/accounts/%s",
		host, accountId.Service, accountId.Account)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodDelete, url, timeout, auth, "")
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode != http.StatusNoContent {
		return res.StatusCode, fmt.Errorf("delete account returned status %d", res.StatusCode)
	}

	return res.StatusCode, nil
}
