package api

import (
	"encoding/json"
	"fmt"
	"mainframe/account/config"
	"mainframe/account/db"
	"net/http"
	"strconv"

	acc "mainframe-lib/account/model"
	com "mainframe-lib/common/config"
	usr "mainframe-lib/user/model"
	susr "mainframe-lib/user/service"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetAccount(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	accId := r.PathValue(string(config.ContextAccount))
	if len(accId) != 24 {
		fmt.Printf("Invalid account id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	serv := r.PathValue(string(config.ContextService))
	if len(serv) != 2 {
		fmt.Printf("Invalid service value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)

	// Build the document
	accountId := acc.AccountId{
		Account: accId,
		Service: serv,
	}

	// Select the document
	account, err := db.SelectAccount(cfg, abi, accountId)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No accounts with id %+v\n", accountId)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching account with id %+v: %s\n", accountId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	serv := r.PathValue(string(config.ContextService))
	if serv != "" && len(serv) != 2 {
		fmt.Printf("Invalid service value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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
	var filter acc.Account
	filter.Id.Account = queryParams.Get(string(config.ContextAccount))
	if serv == "" {
		filter.Id.Service = queryParams.Get(string(config.ContextService))
	} else {
		filter.Id.Service = serv
	}
	filter.Owner = queryParams.Get(string(config.ContextOwner))

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)

	// Select all documents
	accounts, err := db.SelectAccounts(cfg, abi, filter, from, limit)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No accounts with filter %+v\n", filter)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching accounts with filter %+v: %s\n",
			filter, err.Error())
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

func InsertAccount(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var req acc.InsertAccountInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.Id.Account) != 24 {
		fmt.Printf("Invalid account id\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.Id.Service) != 2 {
		fmt.Printf("Invalid account service\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.Owner) != 24 {
		fmt.Printf("Invalid account owner\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)
	auth := r.Context().Value(com.ContextAuth).(string)

	// Build the new document
	account := acc.Account{
		Id: acc.AccountId{
			Account: req.Id.Account,
			Service: req.Id.Service,
		},
		Owner: req.Owner,
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

	// Add account to the user accounts list
	payload := usr.InsertAccountInput{
		Id: usr.AccountId{
			Account: account.Id.Account,
			Service: account.Id.Service,
		},
	}

	err = susr.AddAccountToUser(cfg.Services.Users, cfg.Services.Timeout, auth, account.Owner, payload)
	if err != nil {
		fmt.Printf("Error while adding account %+v to user %s: %s\n",
			account.Id,
			req.Owner,
			err.Error())

		// Rollback
		// Delete the document
		err = db.DeleteAccount(cfg, abi, account.Id)
		if err != nil {
			fmt.Printf("Error while deleting account with id %+v: %s\n",
				account.Id, err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	accId := r.PathValue(string(config.ContextAccount))
	if len(accId) != 24 {
		fmt.Printf("Invalid account id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	serv := r.PathValue(string(config.ContextService))
	if len(serv) != 2 {
		fmt.Printf("Invalid service value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)
	auth := r.Context().Value(com.ContextAuth).(string)

	// Build the document
	accountId := acc.AccountId{
		Account: accId,
		Service: serv,
	}

	// Select the document
	account, err := db.SelectAccount(cfg, abi, accountId)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No accounts with id %s\n", accountId)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching account with id %s: %s\n",
			accountId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Remove account from the user accounts list
	payload := usr.DeleteAccountInput{
		Id: usr.AccountId{
			Account: account.Id.Account,
			Service: account.Id.Service,
		},
	}

	err = susr.RemoveAccountFromUser(cfg.Services.Users, cfg.Services.Timeout, auth, account.Owner, payload)
	if err != nil {
		fmt.Printf("Error while removing account %+v from user %s: %s\n",
			account.Id,
			account.Owner,
			err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Delete the document
	err = db.DeleteAccount(cfg, abi, accountId)
	if err != nil {
		fmt.Printf("Error while deleting account with id %+v: %s\n",
			accountId, err.Error())

		// Rollback
		// Add account to the user accounts list
		payload := usr.InsertAccountInput{
			Id: usr.AccountId{
				Account: account.Id.Account,
				Service: account.Id.Service,
			},
		}

		err = susr.AddAccountToUser(cfg.Services.Users, cfg.Services.Timeout, auth, account.Owner, payload)
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
