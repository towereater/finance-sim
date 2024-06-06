package api

import (
	"net/http"

	"mainframe/user/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete user API function
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Extraction of extra parameters
	id, err := primitive.ObjectIDFromHex(r.PathValue("userId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	// Execution of the request
	err = db.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusNoContent)
}
