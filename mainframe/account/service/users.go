package service

import (
	"errors"
	"fmt"
	"mainframe/account/config"
	"net/http"
)

func AddAccountToUser(cfg config.Config, accountId string, userId string) error {
	// Construct the request
	url := fmt.Sprintf("http://%s/users/%s/accounts/%s", cfg.Services.Users, userId, accountId)

	// Execute the request
	res, err := ExecuteHttpRequest(cfg, http.MethodPost, url, "")
	if err != nil {
		return err
	}

	// Check response
	if res.StatusCode != http.StatusOK {
		return errors.New("Add account to user failed")
	}

	return nil
}

func RemoveAccountFromUser(cfg config.Config, accountId string, userId string) error {
	// Construct the request
	url := fmt.Sprintf("http://%s/users/%s/accounts/%s", cfg.Services.Users, userId, accountId)

	// Execute the request
	res, err := ExecuteHttpRequest(cfg, http.MethodDelete, url, "")
	if err != nil {
		return err
	}

	// Check response
	if res.StatusCode != http.StatusNoContent {
		return errors.New("Delete account from user failed")
	}

	return nil
}
