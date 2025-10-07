package service

import (
	"encoding/json"
	"fmt"
	ccom "mainframe-lib/common/config"
	scom "mainframe-lib/common/service"
	"mainframe-lib/security/model"
	"net/http"
)

func GetUserByApiKey(service ccom.ServiceConfig, auth string, apiKey string) (model.User, int, error) {
	// Construct the request
	url := fmt.Sprintf("/api-keys/%s", apiKey)

	// Execute the request
	res, err := scom.ExecuteHttpRequest(service, http.MethodGet, url, auth, "")
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode == http.StatusNotFound {
		return model.User{}, res.StatusCode, nil
	}
	if res.StatusCode != http.StatusOK {
		return model.User{}, res.StatusCode, fmt.Errorf("get user returned status %d", res.StatusCode)
	}

	// Parse the response
	var user model.User
	err = json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}

	return user, res.StatusCode, nil
}
