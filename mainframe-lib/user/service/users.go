package service

import (
	"encoding/json"
	"errors"
	"fmt"
	com "mainframe-lib/common/service"
	usr "mainframe-lib/user/model"
	"net/http"
)

func GetUser(host string, timeout int, auth string, userId string) (usr.User, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/users/%s", host, userId)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodGet, url, timeout, auth, "")
	if err != nil {
		return usr.User{}, err
	}

	// Check response
	if res.StatusCode != http.StatusOK {
		return usr.User{}, fmt.Errorf("get user returned status %d", res.StatusCode)
	}

	// Parse the response
	var user usr.User
	err = json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		return usr.User{}, err
	}

	return user, nil
}

func GetUserByUsername(host string, timeout int, auth string, username string, password string) (usr.User, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/users?username=%s&password=%s", host, username, password)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodGet, url, timeout, auth, "")
	if err != nil {
		return usr.User{}, err
	}

	// Check response
	if res.StatusCode != http.StatusOK {
		return usr.User{}, fmt.Errorf("get user returned status %d", res.StatusCode)
	}

	// Parse the response
	var users []usr.User
	err = json.NewDecoder(res.Body).Decode(&users)
	if err != nil {
		return usr.User{}, err
	}

	if len(users) > 1 {
		return usr.User{}, fmt.Errorf("get user returned more than one user")
	}

	return users[0], nil
}

func InsertUser(host string, timeout int, auth string, payload usr.InsertUserInput) (usr.User, error) {
	// Construct the request
	url := fmt.Sprintf("http://%s/users", host)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodPost, url, timeout, auth, payload)
	if err != nil {
		return usr.User{}, err
	}

	// Check response
	if res.StatusCode != http.StatusCreated {
		return usr.User{}, fmt.Errorf("insert user returned status %d", res.StatusCode)
	}

	// Parse the response
	var user usr.User
	err = json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		return usr.User{}, err
	}

	return user, nil
}

func UpdateUser(host string, timeout int, auth string, userId string, payload usr.UpdateUserInput) error {
	// Construct the request
	url := fmt.Sprintf("http://%s/users/%s", host, userId)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodPatch, url, timeout, auth, payload)
	if err != nil {
		return err
	}

	// Check response
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("update user returned status %d", res.StatusCode)
	}

	// Parse the response
	var user usr.User
	return json.NewDecoder(res.Body).Decode(&user)
}

func AddAccountToUser(host string, timeout int, auth string, userId string, payload usr.InsertAccountInput) error {
	// Construct the request
	url := fmt.Sprintf("http://%s/users/%s/accounts", host, userId)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodPost, url, timeout, auth, payload)
	if err != nil {
		return err
	}

	// Check response
	if res.StatusCode != http.StatusCreated {
		return errors.New("add account to user failed")
	}

	return nil
}

func RemoveAccountFromUser(host string, timeout int, auth string, userId string, payload usr.DeleteAccountInput) error {
	// Construct the request
	url := fmt.Sprintf("http://%s/users/%s/accounts", host, userId)

	// Execute the request
	res, err := com.ExecuteHttpRequest(http.MethodDelete, url, timeout, auth, payload)
	if err != nil {
		return err
	}

	// Check response
	if res.StatusCode != http.StatusNoContent {
		return errors.New("remove account from user failed")
	}

	return nil
}
