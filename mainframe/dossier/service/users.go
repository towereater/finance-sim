package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"mainframe/dossier/config"
	"mainframe/dossier/model"
	"net/http"
)

func GetUser(cfg config.Config, userId string) (model.User, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/users/%s", cfg.Services.Users, userId)

	// Execute the request
	res, err := ExecuteHttpRequest(cfg, http.MethodGet, url, "", "")
	if err != nil {
		return model.User{}, err
	}

	// Check response
	if res.StatusCode != http.StatusOK {
		return model.User{}, errors.New("get user failed")
	}

	// Parse the response
	var user model.User
	err = json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		return model.User{}, errors.New("get user response convertion failed")
	}

	return user, nil
}
