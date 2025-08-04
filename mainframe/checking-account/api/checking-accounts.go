package api

import (
	"encoding/json"
	"fmt"
	"mainframe/checking-account/config"
	"mainframe/checking-account/db"
	"mainframe/checking-account/model"
	"mainframe/checking-account/service"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCheckingAccount(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	accountId := r.PathValue(string(config.ContextAccountId))
	if accountId == "" || len(accountId) != 24 {
		fmt.Printf("Invalid account id value\n")
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

func GetCheckingAccounts(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	queryParams := r.URL.Query()
	var err error

	from := queryParams.Get(string(config.ContextFrom))
	if from != "" && len(from) != 24 {
		fmt.Printf("Invalid %s parameter\n", string(config.ContextFrom))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	limit := 50
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
	var filter model.CheckingAccount
	filter.IBAN = queryParams.Get("iban")
	filter.Owner = queryParams.Get("owner")

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
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}

func InsertCheckingAccount(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var req model.InsertCheckingAccountInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Owner == "" || len(req.Owner) != 24 {
		fmt.Printf("Invalid account owner\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(config.ContextConfig).(config.Config)
	abi := r.Context().Value(config.ContextAbi).(string)
	cab := r.Context().Value(config.ContextCab).(string)

	// Build the new document
	accountId := primitive.NewObjectID().Hex()
	iban := service.GenerateNewIBAN(abi, cab, accountId)
	account := model.CheckingAccount{
		Id:    accountId,
		Owner: req.Owner,
		Cash:  0,
		IBAN:  iban,
	}

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

	// Insert account into the accounts list
	payload := model.InsertAccountInput{
		Id: model.AccountId{
			Account: account.Id,
			Service: cfg.Prefix,
		},
		Owner: account.Owner,
	}

	err = service.InsertAccount(cfg, payload)
	if err != nil {
		fmt.Printf("Error while adding account %s: %s\n",
			account.Id,
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

func DeleteCheckingAccount(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	accountId := r.PathValue(string(config.ContextAccountId))
	if accountId == "" || len(accountId) != 24 {
		fmt.Printf("Invalid account id value\n")
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

	// Delete account from the accounts list
	err = service.DeleteAccount(cfg, accountId)
	if err != nil {
		fmt.Printf("Error while removing account %s: %s\n",
			accountId,
			err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Delete the document
	err = db.DeleteAccount(cfg, abi, accountId)
	if err != nil {
		fmt.Printf("Error while deleting account with id %s: %s\n", accountId, err.Error())

		// Rollback
		// Insert account to the accounts list
		payload := model.InsertAccountInput{
			Id: model.AccountId{
				Account: account.Id,
				Service: cfg.Prefix,
			},
			Owner: account.Owner,
		}

		err = service.InsertAccount(cfg, payload)
		if err != nil {
			fmt.Printf("Error while adding account %s: %s\n",
				account.Id,
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
