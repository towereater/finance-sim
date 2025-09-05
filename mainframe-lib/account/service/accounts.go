package service

import (
	"encoding/json"
	"errors"
	"fmt"
	acc "mainframe-lib/account/model"
	com "mainframe-lib/common/service"
	"net/http"
)

func GetAccounts(host string, timeout int, auth string, filter acc.Account, from string, limit int) ([]acc.Account, error) {
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
		return []acc.Account{}, err
	}

	// Check response
	if res.StatusCode != http.StatusOK {
		return []acc.Account{}, errors.New("get accounts failed")
	}

	// Parse the response
	var accounts []acc.Account
	err = json.NewDecoder(res.Body).Decode(&accounts)
	if err != nil {
		return []acc.Account{}, err
	}

	return accounts, nil
}

func InsertAccount(host string, timeout int, auth string, payload acc.InsertAccountInput) error {
	// Construct the request
	url := fmt.Sprintf("http://%s/accounts", host)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodPost, url, timeout, auth, payload)
	if err != nil {
		return err
	}

	// Check response
	if res.StatusCode != http.StatusCreated {
		return errors.New("insert account failed")
	}

	return nil
}

func DeleteAccount(host string, timeout int, auth string, accountId acc.AccountId) error {
	// Construct the request
	url := fmt.Sprintf("http://%s/accounts/services/%s/accounts/%s",
		host, accountId.Service, accountId.Account)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodDelete, url, timeout, auth, "")
	if err != nil {
		return err
	}

	// Check response
	if res.StatusCode != http.StatusNoContent {
		return errors.New("delete account failed")
	}

	return nil
}
