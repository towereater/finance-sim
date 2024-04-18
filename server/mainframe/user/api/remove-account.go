package api

import (
	"encoding/json"
	"net/http"

	"mainframe/user/db"
	"mainframe/user/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Remove account user API function
func RemoveAccount(w http.ResponseWriter, r *http.Request, urlModel string) {
	// Parsing of the request
	var req model.RemoveAccountInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extraction of extra parameters
	pathParams := getPathParams(r.URL, urlModel)

	id, err := primitive.ObjectIDFromHex(pathParams["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	// Generation of the data
	accountId, err := primitive.ObjectIDFromHex(req.Account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Execution of the request
	err = db.RemoveAccount(id, accountId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response output
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
