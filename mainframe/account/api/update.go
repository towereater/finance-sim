package api

import (
	"encoding/json"
	"net/http"

	"mainframe/account/db"
	"mainframe/account/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Update account API function
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	// Parsing of the request
	var req model.UpdateAccountInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extraction of extra parameters
	id, err := primitive.ObjectIDFromHex(r.PathValue("accountId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	// Fetch of the account data
	currentAccount, err := db.SelectAccount(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Update account in the user accounts list
	if currentAccount.Owner != req.Owner {
		err = updateAccountOwner(id.Hex(), currentAccount.Owner, req.Owner)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Generation of the updated document
	account := model.Account{
		Id:    id,
		IBAN:  req.IBAN,
		Owner: req.Owner,
		Cash:  req.Cash,
	}

	// Execution of the request
	err = db.UpdateAccount(id, account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response output
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}
