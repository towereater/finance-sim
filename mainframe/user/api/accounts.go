package api

import (
	"encoding/json"
	"fmt"
	"mainframe/user/config"
	"mainframe/user/db"

	com "mainframe-lib/common/config"
	usr "mainframe-lib/user/model"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func AddAccount(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	userId := r.PathValue(string(config.ContextUserId))
	if userId == "" {
		fmt.Printf("Invalid user id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse the request
	var req usr.InsertAccountInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Id.Account == "" || len(req.Id.Account) != 24 {
		fmt.Printf("Invalid account id\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Id.Service == "" || len(req.Id.Service) != 2 {
		fmt.Printf("Invalid account service\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)

	// Build the new document
	account := usr.Account(req)

	// Check user existence
	_, err = db.SelectUser(cfg.DBConfig, abi, userId)
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
	err = db.AddAccount(cfg.DBConfig, abi, userId, account)
	if err != nil {
		fmt.Printf("Error while inserting account %+v to user %s: %s\n",
			account, userId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusCreated)
}

func RemoveAccount(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	userId := r.PathValue(string(config.ContextUserId))
	if userId == "" {
		fmt.Printf("Invalid user id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse the request
	var req usr.DeleteAccountInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Id.Account == "" || len(req.Id.Account) != 24 {
		fmt.Printf("Invalid account id\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Id.Service == "" || len(req.Id.Service) != 2 {
		fmt.Printf("Invalid account service\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)

	// Build the new document
	account := usr.Account(req)

	// Check user existence
	_, err = db.SelectUser(cfg.DBConfig, abi, userId)
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
	err = db.RemoveAccount(cfg.DBConfig, abi, userId, account.Id)
	if err != nil {
		fmt.Printf("Error while removing account %+v from user %s: %s\n",
			account.Id, userId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusNoContent)
}
