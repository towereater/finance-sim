package api

import (
	"encoding/json"
	"fmt"
	"mainframe/dossier/config"
	"mainframe/dossier/db"
	"mainframe/dossier/model"
	"mainframe/dossier/service"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDossier(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	dossierId := r.PathValue(string(config.ContextDossier))
	if dossierId == "" || len(dossierId) != 24 {
		fmt.Printf("Invalid dossier id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(config.ContextConfig).(config.Config)
	abi := r.Context().Value(config.ContextAbi).(string)

	// Select the document
	dossier, err := db.SelectDossier(cfg, abi, dossierId)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No dossiers with id %s\n", dossierId)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching account with id %s: %s\n", dossierId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dossier)
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
	var filter model.Dossier
	filter.Owner = queryParams.Get("owner")

	// Extract context parameters
	cfg := r.Context().Value(config.ContextConfig).(config.Config)
	abi := r.Context().Value(config.ContextAbi).(string)

	// Select all documents
	dossiers, err := db.SelectDossiers(cfg, abi, filter, from, limit)
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
	var req model.InsertDossierInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Owner == "" || len(req.Owner) != 24 {
		fmt.Printf("Invalid dossier owner\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.CheckingAccount.Account == "" || len(req.CheckingAccount.Account) != 24 {
		fmt.Printf("Invalid dossier checking account\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.CheckingAccount.Service == "" || len(req.CheckingAccount.Service) != 2 {
		fmt.Printf("Invalid dossier checking account\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(config.ContextConfig).(config.Config)
	abi := r.Context().Value(config.ContextAbi).(string)

	// Get checking account data
	ckAccount, err := service.GetAccount(cfg, req.CheckingAccount)
	if err != nil {
		fmt.Printf("Error while searching checking account %+v: %s\n",
			req.CheckingAccount, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if ckAccount.Owner != req.Owner {
		fmt.Printf("Checking account owner %s does not match requested owner %s\n",
			ckAccount.Owner, req.CheckingAccount)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get user data
	user, err := service.GetUser(cfg, req.Owner)
	if err != nil {
		fmt.Printf("Error while searching user %s: %s\n", req.Owner, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Build the new document
	dossier := model.Dossier{
		Id:              primitive.NewObjectID().Hex(),
		Owner:           req.Owner,
		CheckingAccount: req.CheckingAccount,
	}

	// Insert dossier into the accounts list
	payload := model.InsertAccountInput{
		Id: model.AccountId{
			Account: dossier.Id,
			Service: cfg.Prefix,
		},
		Owner: dossier.Owner,
	}

	err = service.InsertAccount(cfg, payload)
	if err != nil {
		fmt.Printf("Error while adding dossier %s: %s\n",
			dossier.Id,
			err.Error())

		// // Rollback
		// // Delete the document
		// err = db.DeleteDossier(cfg, abi, dossier.Id)
		// if err != nil {
		// 	fmt.Printf("Error while deleting dossier with id %s: %s\n", dossier.Id, err.Error())

		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a new dossier on xchanger service
	xchangerPayload := model.InsertXChangerDossierInput{
		Name:    user.Name,
		Surname: user.Surname,
		Birth:   user.Birth,
	}

	xchangerDossier, err := service.InsertXChangerDossier(cfg, xchangerPayload)
	if err != nil {
		fmt.Printf("Error while creating xchanger dossier %s: %s\n",
			dossier.Id,
			err.Error())

		// Rollback
		// Delete dossier from the accounts list
		err = db.DeleteDossier(cfg, abi, dossier.Id)
		if err != nil {
			fmt.Printf("Error while deleting dossier with id %s: %s\n", dossier.Id, err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Insert the new document
	dossier.XChangerDossier = xchangerDossier.Id

	err = db.InsertDossier(cfg, abi, dossier)
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

	// Response output
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dossier)
}

func DeleteDossier(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	dossierId := r.PathValue(string(config.ContextDossier))
	if dossierId == "" || len(dossierId) != 24 {
		fmt.Printf("Invalid dossier id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(config.ContextConfig).(config.Config)
	abi := r.Context().Value(config.ContextAbi).(string)

	// Select the document
	dossier, err := db.SelectDossier(cfg, abi, dossierId)
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

	// Delete dossier from xchanger service
	err = service.DeleteXChangerDossier(cfg, dossier.XChangerDossier)
	if err != nil {
		fmt.Printf("Error while deleting xchanger dossier %s: %s\n",
			dossier.Id,
			err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Delete dossier from the accounts list
	err = service.DeleteAccount(cfg, dossierId)
	if err != nil {
		fmt.Printf("Error while removing dossier %s: %s\n",
			dossierId,
			err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Delete the document
	err = db.DeleteDossier(cfg, abi, dossierId)
	if err != nil {
		fmt.Printf("Error while deleting dossier with id %s: %s\n", dossierId, err.Error())

		// Rollback
		// Insert dossier to the accounts list
		payload := model.InsertAccountInput{
			Id: model.AccountId{
				Account: dossier.Id,
				Service: cfg.Prefix,
			},
			Owner: dossier.Owner,
		}

		err = service.InsertAccount(cfg, payload)
		if err != nil {
			fmt.Printf("Error while adding dossier %s: %s\n",
				dossier.Id,
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
