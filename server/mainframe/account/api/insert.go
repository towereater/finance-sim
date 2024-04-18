package api

import (
	"encoding/json"
	"net/http"

	"mainframe/account/db"
	"mainframe/account/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert account API function
func InsertAccount(w http.ResponseWriter, r *http.Request) {
	// Parsing of the request
	var req model.InsertAccountInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generation of the new document
	account := model.Account{
		Id:    primitive.NewObjectID(),
		IBAN:  primitive.NewObjectID().Hex(),
		Owner: req.Owner,
		Cash:  0,
	}

	// Execution of the request
	err = db.InsertAccount(account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response output
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}
