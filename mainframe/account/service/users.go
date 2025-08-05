package service

import (
	"errors"
	"fmt"
	"mainframe/account/config"
	"mainframe/account/model"
	"net/http"
)

func AddAccountToUser(cfg config.Config, auth string, userId string, payload model.AddAccountToUserInput) error {
	// Construct the request
	url := fmt.Sprintf("http://%s/users/%s/accounts", cfg.Services.Users, userId)

	// Execute the request
	res, err := ExecuteHttpRequest(cfg, http.MethodPost, url, auth, payload)
	if err != nil {
		return err
	}

	// Check response
	if res.StatusCode != http.StatusOK {
		return errors.New("add account to user failed")
	}

	return nil
}

func RemoveAccountFromUser(cfg config.Config, auth string, userId string, payload model.RemoveAccountFromUserInput) error {
	// Construct the request
	url := fmt.Sprintf("http://%s/users/%s/accounts", cfg.Services.Users, userId)

	// Execute the request
	res, err := ExecuteHttpRequest(cfg, http.MethodDelete, url, auth, payload)
	if err != nil {
		return err
	}

	// Check response
	if res.StatusCode != http.StatusNoContent {
		return errors.New("remove account from user failed")
	}

	return nil
}
