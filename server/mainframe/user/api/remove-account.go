package api

import (
	"net/http"

	"mainframe/user/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Remove account user API function
func RemoveAccount(w http.ResponseWriter, r *http.Request) {
	// Extraction of extra parameters
	id, err := primitive.ObjectIDFromHex(r.PathValue("userId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	accountId, err := primitive.ObjectIDFromHex(r.PathValue("accountId"))
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
