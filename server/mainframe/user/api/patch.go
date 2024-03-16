package api

import (
	"encoding/json"
	"net/http"

	"mainframe/user/db"
	"mainframe/user/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Patch user API function
func PatchUser(w http.ResponseWriter, r *http.Request, urlModel string) {
	// Parsing of the request
	var req model.PatchUserInput
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

	// Fetch of the user data
	user, err := db.SelectUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generation of the updated document
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Password != "" {
		user.Password = req.Password
	}
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Surname != "" {
		user.Surname = req.Surname
	}
	if req.Birth != "" {
		user.Birth = req.Birth
	}

	// Execution of the request
	err = db.UpdateUser(id, *user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response output
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
