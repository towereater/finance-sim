package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"mainframe/account/db"
	"mainframe/account/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get account API function
func GetAccount(w http.ResponseWriter, r *http.Request) {
	// Extraction of extra parameters
	id, err := primitive.ObjectIDFromHex(r.PathValue("accountId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	// Execution of the request
	user, err := db.SelectAccount(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response output
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// Get accounts API function
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	// Extraction of extra parameters
	queryParams := r.URL.Query()

	fromId := primitive.NilObjectID
	if queryParams.Has("fromId") {
		fromId, _ = primitive.ObjectIDFromHex(queryParams.Get("fromId"))
	}

	limit := 0
	if queryParams.Has("limit") {
		limit, _ = strconv.Atoi(queryParams.Get("limit"))
	}

	order := 1
	if queryParams.Has("order") {
		order, _ = strconv.Atoi(queryParams.Get("order"))
	}

	// Building the filter
	var filter model.Account

	if queryParams.Has("iban") {
		filter.IBAN = queryParams.Get("iban")
	}

	if queryParams.Has("owner") {
		filter.Owner = queryParams.Get("owner")
	}

	// Execution of the request
	accounts, err := db.SelectAccounts(filter, fromId, limit, order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if accounts == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Response output
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}
