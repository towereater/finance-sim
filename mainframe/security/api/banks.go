package api

import (
	"encoding/json"
	"fmt"
	com "mainframe-lib/common/config"
	sec "mainframe-lib/security/model"
	"mainframe/security/config"
	"mainframe/security/db"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetBank(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	abi := r.PathValue(string(config.ContextAbi))
	if len(abi) != 5 {
		fmt.Printf("Invalid abi value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)

	// Select the document
	bank, err := db.SelectBankByAbi(cfg.DBConfig, abi)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No banks with abi %s\n", abi)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching bank with abi %s: %s\n", abi, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bank)
}

func InsertBank(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var req sec.InsertBankInput
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

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)

	// Build the new document
	bank := sec.Bank(req)

	// Insert the new document
	err = db.InsertBank(cfg.DBConfig, bank)
	if mongo.IsDuplicateKeyError(err) {
		fmt.Printf("Bank %+v already exists\n", bank)
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err != nil {
		fmt.Printf("Error while inserting bank %+v: %s\n", bank, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bank)
}

func DeleteBank(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	abi := r.PathValue(string(config.ContextAbi))
	if len(abi) != 5 {
		fmt.Printf("Invalid abi value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)

	// Delete the document
	err := db.DeleteBank(cfg.DBConfig, abi)
	if err != nil {
		fmt.Printf("Error while deleting bank with id %s: %s\n", abi, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusNoContent)
}
