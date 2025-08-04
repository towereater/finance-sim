package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"mainframe/security/config"
	"mainframe/security/db"
	"mainframe/security/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetApiKey(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	apiKeyId := r.PathValue(string(config.ContextApiKey))
	if len(apiKeyId) != 24 {
		fmt.Printf("Invalid api key id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(config.ContextConfig).(config.Config)
	abi := r.Context().Value(config.ContextAbi).(string)

	// Select the document
	apiKey, err := db.SelectApiKey(cfg, abi, apiKeyId)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No api keys with id %s\n", apiKeyId)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching api key with id %s: %s\n", apiKeyId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiKey)
}

func InsertApiKey(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var req model.InsertApiKeyInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.Abi) != 5 {
		fmt.Printf("Invalid abi\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.Cab) != 5 {
		fmt.Printf("Invalid cab\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(config.ContextConfig).(config.Config)
	abi := r.Context().Value(config.ContextAbi).(string)

	// Build the new document
	apiKey := model.ApiKey{
		Id:  primitive.NewObjectID().Hex(),
		Abi: req.Abi,
		Cab: req.Cab,
	}

	// Insert the new document
	err = db.InsertApiKey(cfg, abi, apiKey)
	if mongo.IsDuplicateKeyError(err) {
		fmt.Printf("Api key %+v already exists\n", apiKey)
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err != nil {
		fmt.Printf("Error while inserting api key %+v: %s\n", apiKey, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(apiKey)
}

func DeleteApiKey(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	apiKeyId := r.PathValue(string(config.ContextApiKey))
	if len(apiKeyId) != 24 {
		fmt.Printf("Invalid api key id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(config.ContextConfig).(config.Config)
	abi := r.Context().Value(config.ContextAbi).(string)

	// Delete the document
	err := db.DeleteApiKey(cfg, abi, apiKeyId)
	if err != nil {
		fmt.Printf("Error while deleting api key with id %s: %s\n", apiKeyId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusNoContent)
}
