package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"mainframe/dossier/config"
	"mainframe/dossier/model"
	"net/http"
)

func GetAccount(cfg config.Config, accountId model.AccountId) (model.Account, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/accounts/services/%s/accounts/%s",
		cfg.Services.Accounts, accountId.Service, accountId.Account)

	// Execute the request
	res, err := ExecuteHttpRequest(cfg, http.MethodGet, url, "", "")
	if err != nil {
		return model.Account{}, err
	}

	// Check response
	if res.StatusCode != http.StatusOK {
		return model.Account{}, errors.New("get account failed")
	}

	// Parse the response
	var account model.Account
	err = json.NewDecoder(res.Body).Decode(&account)
	if err != nil {
		return model.Account{}, errors.New("get account response convertion failed")
	}

	return account, nil
}

func InsertAccount(cfg config.Config, payload model.InsertAccountInput) error {
	// Construct the request
	url := fmt.Sprintf("http://%s/accounts", cfg.Services.Accounts)

	// Execute the request
	res, err := ExecuteHttpRequest(cfg, http.MethodPost, url, payload, "")
	if err != nil {
		return err
	}

	// Check response
	if res.StatusCode != http.StatusCreated {
		return errors.New("insert account failed")
	}

	return nil
}

func DeleteAccount(cfg config.Config, accountId string) error {
	// Construct the request
	url := fmt.Sprintf("http://%s/accounts/services/%s/accounts/%s",
		cfg.Services.Accounts, cfg.Prefix, accountId)

	// Execute the request
	res, err := ExecuteHttpRequest(cfg, http.MethodDelete, url, "", "")
	if err != nil {
		return err
	}

	// Check response
	if res.StatusCode != http.StatusNoContent {
		return errors.New("delete account failed")
	}

	return nil
}
