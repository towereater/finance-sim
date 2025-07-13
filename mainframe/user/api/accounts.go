package api

import (
	"fmt"
	"mainframe/user/config"
	"mainframe/user/db"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// Insert the new document
	err = db.AddAccount(cfg, abi, userId, accountId)
	if err != nil {
		fmt.Printf("Error while inserting account %s to user %s: %s\n",
			accountId, userId, err.Error())
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
