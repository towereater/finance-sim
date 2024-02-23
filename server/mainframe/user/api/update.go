package api

import (
	"encoding/json"
	"net/http"

	"mainframe/user/db"
	"mainframe/user/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Update user API function
func UpdateUser(w http.ResponseWriter, r *http.Request, urlModel string) {
	// Parsing of the request
	var req model.UpdateUserInput
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

	// Generation of the updated document
	user := model.User{
		Id:       id,
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		Surname:  req.Surname,
	}

	// Execution of the request
	err = db.UpdateUser(id, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
