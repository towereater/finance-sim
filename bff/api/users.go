package api

import (
	"bff/config"
	"bff/model"
	"bff/service"
	"encoding/json"
	"fmt"
	"net/http"

	com "mainframe-lib/common/config"
	usr "mainframe-lib/user/model"
	susr "mainframe-lib/user/service"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var req usr.InsertUserInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	auth := r.Context().Value(com.ContextAuth).(string)

	// Create a new user
	_, err = susr.InsertUser(cfg.Services.Users, cfg.Services.Timeout, auth, req)
	if err != nil {
		fmt.Printf("Error while creating user: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Response output
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var req model.GetUserInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	auth := r.Context().Value(com.ContextAuth).(string)

	// Create a new user
	user, err := susr.GetUserByUsername(cfg.Services.Users, cfg.Services.Timeout, auth, req.Username, req.Password)
	if err != nil {
		fmt.Printf("Error while getting user: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create a JWT token
	jwt := service.CreateJWT(user)

	// Create response output
	res := model.GetUserOutput{
		Name:          user.Name,
		Surname:       user.Surname,
		Authorization: jwt,
	}

	// Response output
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func ChangeUserPassword(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var req usr.UpdateUserInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	auth := r.Context().Value(com.ContextAuth).(string)
	userId := r.Context().Value(config.ContextUserId).(string)

	// Create a new user
	err = susr.UpdateUser(cfg.Services.Users, cfg.Services.Timeout, auth, userId, req)
	if err != nil {
		fmt.Printf("Error while updating user: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Response output
	w.WriteHeader(http.StatusOK)
}

func ResetUserPassword(w http.ResponseWriter, r *http.Request) {
	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	auth := r.Context().Value(com.ContextAuth).(string)
	userId := r.Context().Value(config.ContextUserId).(string)

	// Update user password
	req := usr.UpdateUserInput{
		Password: "password",
	}

	// Create a new user
	err := susr.UpdateUser(cfg.Services.Users, cfg.Services.Timeout, auth, userId, req)
	if err != nil {
		fmt.Printf("Error while updating user: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Response output
	w.WriteHeader(http.StatusOK)
}
