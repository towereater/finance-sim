package api

import (
	"encoding/json"
	"fmt"
	"mainframe/account/config"
	"mainframe/account/db"
	"mainframe/account/model"
	"mainframe/account/service"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAccount(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	id := r.PathValue(string(config.ContextAccountId))
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

	// Select the document
	account, err := db.SelectAccount(cfg, abi, accountId)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No accounts with id %s\n", accountId)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching account with id %s: %s\n", accountId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	queryParams := r.URL.Query()
	var err error

	from := primitive.NilObjectID
	if queryParams.Has(string(config.ContextFrom)) {
		from, err = primitive.ObjectIDFromHex(queryParams.Get(string(config.ContextFrom)))

		if err != nil {
			fmt.Printf("Invalid %s parameter\n", string(config.ContextFrom))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	limit := 0
	if queryParams.Has(string(config.ContextLimit)) {
		limit, err = strconv.Atoi(queryParams.Get(string(config.ContextLimit)))

		if err != nil {
			fmt.Printf("Invalid %s parameter\n", string(config.ContextLimit))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if limit > 50 {
			limit = 50
		}
	}

	// Build the filter
	var filter model.Account

	if queryParams.Has("owner") {
		filter.Owner = queryParams.Get("owner")
	}
	if queryParams.Has("service") {
		filter.Owner = queryParams.Get("service")
	}

	// Extract context parameters
	cfg := r.Context().Value(config.ContextConfig).(config.Config)
	abi := r.Context().Value(config.ContextAbi).(string)

	// Select all documents
	accounts, err := db.SelectAccounts(cfg, abi, filter, from, limit)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No accounts with filter %+v\n", filter)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching accounts with filter %+v: %s\n", filter, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	if len(accounts) == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}

func InsertAccount(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var req model.InsertAccountInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Owner == "" {
		fmt.Printf("Invalid account owner\n")
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
		Id:      primitive.NewObjectID(),
		Owner:   req.Owner,
		Service: req.Service,
	}

	// Extract context parameters
	cfg := r.Context().Value(config.ContextConfig).(config.Config)
	abi := r.Context().Value(config.ContextAbi).(string)

	// Insert the new document
	err = db.InsertAccount(cfg, abi, account)
	if mongo.IsDuplicateKeyError(err) {
		fmt.Printf("Account %+v already exists\n", account)
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err != nil {
		fmt.Printf("Error while inserting account %+v: %s\n", account, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Add account to the user accounts list
	payload := model.AddAccountToUserInput{
		Id:      account.Id,
		Service: account.Service,
	}

	err = service.AddAccountToUser(cfg, account.Owner, payload)
	if err != nil {
		fmt.Printf("Error while adding account %s to user %s: %s\n",
			account.Id,
			req.Owner,
			err.Error())

		// Rollback
		// Delete the document
		err = db.DeleteAccount(cfg, abi, account.Id)
		if err != nil {
			fmt.Printf("Error while deleting account with id %s: %s\n", account.Id, err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	id := r.PathValue(string(config.ContextAccountId))
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

	// Select the document
	account, err := db.SelectAccount(cfg, abi, accountId)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No accounts with id %s\n", accountId)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching account with id %s: %s\n", accountId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Remove account from the user accounts list
	err = service.RemoveAccountFromUser(cfg, account.Owner, accountId.Hex())
	if err != nil {
		fmt.Printf("Error while removing account %s from user %s: %s\n",
			account.Id,
			account.Owner,
			err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Delete the document
	err = db.DeleteAccount(cfg, abi, accountId)
	if err != nil {
		fmt.Printf("Error while deleting account with id %s: %s\n", accountId, err.Error())

		// Rollback
		// Add account to the user accounts list
		payload := model.AddAccountToUserInput{
			Id:      account.Id,
			Service: account.Service,
		}

		err = service.AddAccountToUser(cfg, account.Owner, payload)
		if err != nil {
			fmt.Printf("Error while adding account %s to user %s: %s\n",
				account.Id,
				account.Owner,
				err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusNoContent)
}
