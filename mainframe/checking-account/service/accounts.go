package service

import (
	"errors"
	"fmt"
	"mainframe/checking-account/config"
	"mainframe/checking-account/model"
	"net/http"
)

func InsertAccount(cfg config.Config, auth string, payload model.InsertAccountInput) error {
	// Construct the request
	url := fmt.Sprintf("http://%s/accounts", cfg.Services.Accounts)

	// Execute the request
	res, err := ExecuteHttpRequest(cfg, http.MethodPost, url, auth, payload)
	if err != nil {
		return err
	}

	// Check response
	if res.StatusCode != http.StatusCreated {
		return errors.New("insert account failed")
	}

	return nil
}

func DeleteAccount(cfg config.Config, auth string, accountId string) error {
	// Construct the request
	url := fmt.Sprintf("http://%s/accounts/services/%s/accounts/%s",
		cfg.Services.Accounts, cfg.Prefix, accountId)

	// Execute the request
	res, err := ExecuteHttpRequest(cfg, http.MethodDelete, url, auth, "")
	if err != nil {
		return err
	}

	// Check response
	if res.StatusCode != http.StatusNoContent {
		return errors.New("delete account failed")
	}

	return nil
}
