package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	com "mainframe-lib/common/config"
	"mainframe/security/config"
	"mainframe/security/db"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetApiKey(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	apiKey := r.PathValue(string(config.ContextApiKey))
	if len(apiKey) != 24 {
		fmt.Printf("Invalid api key value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)

	// Select the document
	user, err := db.SelectUserByApiKey(cfg.DBConfig, abi, apiKey)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No users with api key %s\n", apiKey)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching user with api key %s: %s\n", apiKey, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusOK)
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(user)
	}
}
