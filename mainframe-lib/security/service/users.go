package service

import (
	"encoding/json"
	"fmt"
	com "mainframe-lib/common/service"
	"mainframe-lib/security/model"
	"net/http"
)

func GetUserByApiKey(host string, timeout int, auth string, apiKey string) (model.User, int, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/api-keys/%s", host, apiKey)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodGet, url, timeout, auth, "")
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
