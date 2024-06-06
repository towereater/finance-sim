package api

import (
	"encoding/json"
	"net/http"

	"mainframe/user/db"
	"mainframe/user/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Update user API function
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Parsing of the request
	var req model.UpdateUserInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extraction of extra parameters
	id, err := primitive.ObjectIDFromHex(r.PathValue("userId"))
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
		Birth:    req.Birth,
		Accounts: req.Accounts,
	}
	if req.Accounts == nil {
		user.Accounts = make([]string, 0)
	}

	// Execution of the request
	err = db.UpdateUser(id, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response output
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
