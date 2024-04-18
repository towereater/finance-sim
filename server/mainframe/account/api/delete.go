package api

import (
	"net/http"

	"mainframe/user/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete user API function
func DeleteUser(w http.ResponseWriter, r *http.Request, urlModel string) {
	// Extraction of extra parameters
	pathParams := getPathParams(r.URL, urlModel)

	id, err := primitive.ObjectIDFromHex(pathParams["id"])
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
