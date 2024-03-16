package api

import (
	"encoding/json"
	"net/http"

	"mainframe/user/db"
	"mainframe/user/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert user API function
func InsertUser(w http.ResponseWriter, r *http.Request) {
	// Parsing of the request
	var req model.InsertUserInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generation of the new document
	user := model.User{
		Id:       primitive.NewObjectID(),
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		Surname:  req.Surname,
		Birth:    req.Birth,
	}

	// Execution of the request
	err = db.InsertUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response output
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
