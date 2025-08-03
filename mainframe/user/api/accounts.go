package api

import (
	"encoding/json"
	"fmt"
	"mainframe/user/config"
	"mainframe/user/db"
	"mainframe/user/model"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddAccount(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	id := r.PathValue(string(config.ContextUserId))
	if id == "" {
		fmt.Printf("Invalid user id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error while converting user id %s: %s\n", id, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse the request
	var req model.InsertAccountInput
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Id.Hex() == "" {
		fmt.Printf("Invalid account id\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Service == "" || len(req.Service) != 2 {
		fmt.Printf("Invalid account service\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Build the new document
	account := model.Account{
		Id:      req.Id,
		Service: req.Service,
	}

	// Extract context parameters
	cfg := r.Context().Value(config.ContextConfig).(config.Config)
	abi := r.Context().Value(config.ContextAbi).(string)

	// Check user existence
	_, err = db.SelectUser(cfg, abi, userId)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No users with id %s\n", userId)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching user with id %s: %s\n", userId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Insert the new document
	err = db.AddAccount(cfg, abi, userId, account)
	if err != nil {
		fmt.Printf("Error while inserting account %+v to user %s: %s\n",
			account, userId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusOK)
}

func RemoveAccount(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	id := r.PathValue(string(config.ContextUserId))
	if id == "" {
		fmt.Printf("Invalid user id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error while converting user id %s: %s\n", id, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id = r.PathValue(string(config.ContextAccountId))
	if id == "" {
		fmt.Printf("Invalid account id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accountId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error while converting account id %s: %s\n", id, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(config.ContextConfig).(config.Config)
	abi := r.Context().Value(config.ContextAbi).(string)

	// Check user existence
	_, err = db.SelectUser(cfg, abi, userId)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No users with id %s\n", userId)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching user with id %s: %s\n", userId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Remove the document
	err = db.RemoveAccount(cfg, abi, userId, accountId)
	if err != nil {
		fmt.Printf("Error while removing account %s from user %s: %s\n",
			accountId, userId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusNoContent)
}
