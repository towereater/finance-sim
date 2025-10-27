package api

import (
	"encoding/json"
	"fmt"
	"mainframe/dossier/config"
	"mainframe/dossier/db"
	"mainframe/dossier/service"
	"net/http"
	"strconv"

	acc "mainframe-lib/account/model"
	sacc "mainframe-lib/account/service"
	scha "mainframe-lib/checking-account/service"
	com "mainframe-lib/common/config"
	dos "mainframe-lib/dossier/model"
	ssec "mainframe-lib/security/service"
	susr "mainframe-lib/user/service"
	xch "mainframe-lib/xchanger/model"
	sxch "mainframe-lib/xchanger/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDossier(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	dossierId := r.PathValue(string(config.ContextDossierId))
	if len(dossierId) != 24 {
		fmt.Printf("Invalid dossier id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)
	auth := r.Context().Value(com.ContextAuth).(string)

	// Select the document
	dossier, err := db.SelectDossier(cfg.DB, abi, dossierId)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No dossiers with id %s\n", dossierId)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching dossier with id %s: %s\n", dossierId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get bank data
	bank, status, err := ssec.GetBankByAbi(cfg.Services.Security, auth, abi)
	if err != nil {
		fmt.Printf("Error while searching bank with abi %s: %s\n", abi, err.Error())
		w.WriteHeader(status)
		return
	}
	if bank.XchangerApiKey == "" {
		fmt.Printf("Bank with abi %s does not have access to xchanger\n", abi)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Select the document
	xchangerDossier, status, err := sxch.GetDossier(
		cfg.Services.Xchanger, bank.XchangerApiKey, dossier.XChangerDossier)
	if err != nil {
		fmt.Printf("Error while searching dossier %s on xchanger: %s\n",
			dossier.XChangerDossier,
			err.Error())
		w.WriteHeader(status)
		return
	}

	// Convert the document to standard format
	dossierDto := service.ToDossierDto(dossier, xchangerDossier)

	// Response output
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dossierDto)
}

func GetDossiers(w http.ResponseWriter, r *http.Request) {
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
	var filter dos.Dossier
	filter.Owner = queryParams.Get("owner")

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)

	// Select all documents
	dossiers, err := db.SelectDossiers(cfg.DB, abi, filter, from, limit)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No dossiers with filter %+v\n", filter)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching dossiers with filter %+v: %s\n", filter, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	if len(dossiers) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dossiers)
}

func InsertDossier(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var req dos.InsertDossierInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.Owner) != 24 {
		fmt.Printf("Invalid dossier owner\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.CheckingAccount.Account) != 24 {
		fmt.Printf("Invalid dossier checking account\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.CheckingAccount.Service) != 2 {
		fmt.Printf("Invalid dossier checking account\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)
	auth := r.Context().Value(com.ContextAuth).(string)

	// Get bank data
	bank, status, err := ssec.GetBankByAbi(cfg.Services.Security, auth, abi)
	if err != nil {
		fmt.Printf("Error while searching bank with abi %s: %s\n", abi, err.Error())
		w.WriteHeader(status)
		return
	}
	if bank.XchangerApiKey == "" {
		fmt.Printf("Bank with abi %s does not have access to xchanger\n", abi)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Get checking account data
	ckAccount, status, err := scha.GetAccount(cfg.Services.CheckingAccounts, auth, req.CheckingAccount.Account)
	if err != nil {
		fmt.Printf("Error while searching checking account %+v: %s\n",
			req.CheckingAccount, err.Error())
		w.WriteHeader(status)
		return
	}
	if ckAccount.Owner != req.Owner {
		fmt.Printf("Checking account owner %s does not match requested owner %s\n",
			ckAccount.Owner, req.CheckingAccount)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get user data
	user, status, err := susr.GetUser(cfg.Services.Users, auth, req.Owner)
	if err != nil {
		fmt.Printf("Error while searching user %s: %s\n", req.Owner, err.Error())
		w.WriteHeader(status)
		return
	}

	// Build the new document
	dossier := dos.Dossier{
		Id:              primitive.NewObjectID().Hex(),
		Owner:           req.Owner,
		CheckingAccount: req.CheckingAccount,
	}

	// Insert dossier into the accounts list
	payload := acc.InsertAccountInput{
		Id: acc.AccountId{
			Account: dossier.Id,
			Service: cfg.DB.Prefix,
		},
		Owner: dossier.Owner,
	}

	status, err = sacc.InsertAccount(cfg.Services.Accounts, auth, payload)
	if err != nil {
		fmt.Printf("Error while adding dossier %s: %s\n",
			dossier.Id,
			err.Error())

		w.WriteHeader(status)
		return
	}

	// Create a new dossier on xchanger
	xchangerPayload := xch.InsertDossierInput{
		Name:       user.Name,
		Surname:    user.Surname,
		Birth:      user.Birth,
		ExternalId: dossier.Id,
		IBAN:       ckAccount.IBAN,
	}

	xchangerDossier, status, err := sxch.InsertDossier(cfg.Services.Xchanger, bank.XchangerApiKey, xchangerPayload)
	if err != nil {
		fmt.Printf("Error while creating xchanger dossier %s: %s\n",
			dossier.Id,
			err.Error())

		// Rollback
		// Delete dossier from the accounts list
		err = db.DeleteDossier(cfg.DB, abi, dossier.Id)
		if err != nil {
			fmt.Printf("Error while deleting dossier with id %s: %s\n", dossier.Id, err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(status)
		return
	}

	// Insert the new document
	dossier.XChangerDossier = xchangerDossier.Id

	err = db.InsertDossier(cfg.DB, abi, dossier)
	if mongo.IsDuplicateKeyError(err) {
		fmt.Printf("Dossier %+v already exists\n", dossier)
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err != nil {
		fmt.Printf("Error while inserting dossier %+v: %s\n", dossier, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Convert the document to standard format
	dossierDto := service.ToDossierDto(dossier, xchangerDossier)

	// Response output
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dossierDto)
}

func DeleteDossier(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	dossierId := r.PathValue(string(config.ContextDossierId))
	if len(dossierId) != 24 {
		fmt.Printf("Invalid dossier id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)
	auth := r.Context().Value(com.ContextAuth).(string)

	// Select the document
	dossier, err := db.SelectDossier(cfg.DB, abi, dossierId)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No dossiers with id %s\n", dossierId)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching dossier with id %s: %s\n", dossierId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get bank data
	bank, status, err := ssec.GetBankByAbi(cfg.Services.Security, auth, abi)
	if err != nil {
		fmt.Printf("Error while searching bank with abi %s: %s\n", abi, err.Error())
		w.WriteHeader(status)
		return
	}
	if bank.XchangerApiKey == "" {
		fmt.Printf("Bank with abi %s does not have access to xchanger\n", abi)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Delete dossier from xchanger
	status, err = sxch.DeleteDossier(cfg.Services.Xchanger, bank.XchangerApiKey, dossier.XChangerDossier)
	if err != nil {
		fmt.Printf("Error while deleting xchanger dossier %s: %s\n",
			dossier.Id,
			err.Error())

		w.WriteHeader(status)
		return
	}

	// Delete dossier from the accounts list
	accountId := acc.AccountId{
		Account: dossierId,
		Service: "DS",
	}
	status, err = sacc.DeleteAccount(cfg.Services.Accounts, auth, accountId)
	if err != nil {
		fmt.Printf("Error while removing dossier %s: %s\n",
			dossierId,
			err.Error())
		w.WriteHeader(status)
		return
	}

	// Delete the document
	err = db.DeleteDossier(cfg.DB, abi, dossierId)
	if err != nil {
		fmt.Printf("Error while deleting dossier with id %s: %s\n", dossierId, err.Error())

		// Rollback
		// Insert dossier to the accounts list
		payload := acc.InsertAccountInput{
			Id: acc.AccountId{
				Account: dossier.Id,
				Service: cfg.DB.Prefix,
			},
			Owner: dossier.Owner,
		}

		status, err = sacc.InsertAccount(cfg.Services.Accounts, auth, payload)
		if err != nil {
			fmt.Printf("Error while adding dossier %s: %s\n",
				dossier.Id,
				err.Error())
			w.WriteHeader(status)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusNoContent)
}
