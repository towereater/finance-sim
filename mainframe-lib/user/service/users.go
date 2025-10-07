package service

import (
	"encoding/json"
	"fmt"
	ccom "mainframe-lib/common/config"
	scom "mainframe-lib/common/service"
	"mainframe-lib/user/model"
	"net/http"
)

func GetUser(service ccom.ServiceConfig, auth string, userId string) (model.User, int, error) {
	// Construct the request
	url := fmt.Sprintf("/users/%s", userId)

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

func GetUserByUsername(service ccom.ServiceConfig, auth string, username string, password string) (model.User, int, error) {
	// Construct the request
	url := fmt.Sprintf("/users?username=%s&password=%s", username, password)

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
	var users []model.User
	err = json.NewDecoder(res.Body).Decode(&users)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}

	if len(users) > 1 {
		return model.User{}, http.StatusInternalServerError, fmt.Errorf("get user returned more than one user")
	}

	return users[0], res.StatusCode, nil
}

func InsertUser(service ccom.ServiceConfig, auth string, payload model.InsertUserInput) (model.User, int, error) {
	// Construct the request
	url := "/users"

	// Execute the request
	res, err := scom.ExecuteHttpRequest(service, http.MethodPost, url, auth, payload)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode != http.StatusCreated {
		return model.User{}, res.StatusCode, fmt.Errorf("insert user returned status %d", res.StatusCode)
	}

	// Parse the response
	var user model.User
	err = json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}

	return user, res.StatusCode, nil
}

func UpdateUser(service ccom.ServiceConfig, auth string, userId string, payload model.UpdateUserInput) (model.User, int, error) {
	// Construct the request
	url := fmt.Sprintf("/users/%s", userId)

	// Execute the request
	res, err := scom.ExecuteHttpRequest(service, http.MethodPatch, url, auth, payload)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode != http.StatusOK {
		return model.User{}, res.StatusCode, fmt.Errorf("update user returned status %d", res.StatusCode)
	}

	// Parse the response
	var user model.User
	err = json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}

	return user, res.StatusCode, nil
}

func AddAccountToUser(service ccom.ServiceConfig, auth string, userId string, payload model.InsertAccountInput) (int, error) {
	// Construct the request
	url := fmt.Sprintf("/users/%s/accounts", userId)

	// Execute the request
	res, err := scom.ExecuteHttpRequest(service, http.MethodPost, url, auth, payload)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode != http.StatusCreated {
		return res.StatusCode, fmt.Errorf("add account to user failed")
	}

	return res.StatusCode, nil
}

func RemoveAccountFromUser(service ccom.ServiceConfig, auth string, userId string, payload model.DeleteAccountInput) (int, error) {
	// Construct the request
	url := fmt.Sprintf("/users/%s/accounts", userId)

	// Execute the request
	res, err := scom.ExecuteHttpRequest(service, http.MethodDelete, url, auth, payload)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Check response
	if res.StatusCode != http.StatusNoContent {
		return res.StatusCode, fmt.Errorf("remove account from user failed")
	}

	return res.StatusCode, nil
}
