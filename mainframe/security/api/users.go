package api

import (
	"encoding/json"
	"fmt"
	com "mainframe-lib/common/config"
	sec "mainframe-lib/security/model"
	"mainframe/security/config"
	"mainframe/security/db"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var req sec.InsertUserInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.Username) < 1 {
		fmt.Printf("Invalid username\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.Password) < 8 {
		fmt.Printf("Invalid password\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.ApiKey) != 24 {
		fmt.Printf("Invalid api key\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)

	// Build the new document
	user := sec.User{
		Id:       primitive.NewObjectID().Hex(),
		Username: req.Username,
		Password: req.Password,
		ApiKey:   req.ApiKey,
	}

	// Insert the new document
	err = db.InsertUser(cfg, abi, user)
	if mongo.IsDuplicateKeyError(err) {
		fmt.Printf("User %+v already exists\n", user)
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err != nil {
		fmt.Printf("Error while inserting user %+v: %s\n", user, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	userId := r.PathValue(string(config.ContextUserId))
	if len(userId) != 24 {
		fmt.Printf("Invalid user id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)

	// Delete the document
	err := db.DeleteUser(cfg, abi, userId)
	if err != nil {
		fmt.Printf("Error while deleting user with id %s: %s\n", userId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusNoContent)
}
