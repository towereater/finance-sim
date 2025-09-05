package service

import (
	"errors"
	"fmt"
	acc "mainframe-lib/account/model"
	com "mainframe-lib/common/service"
	"net/http"
)

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
