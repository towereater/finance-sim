package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"mainframe/user/config"
	"mainframe/user/db"
	"mainframe/user/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	id := r.PathValue(string(config.ContextUserId))
	if id == "" {
		fmt.Printf("Invalid user id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error while converting user id %s: %s\n", id, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(config.ContextConfig).(config.Config)
	abi := r.Context().Value(config.ContextAbi).(string)

	// Select the document
	user, err := db.SelectUser(cfg, abi, userId)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No users with id %s\n", userId)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching user with id %s: %s\n", userId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	queryParams := r.URL.Query()
	var err error

	from := primitive.NilObjectID
	if queryParams.Has(string(config.ContextFrom)) {
		from, err = primitive.ObjectIDFromHex(queryParams.Get(string(config.ContextFrom)))

		if err != nil {
			fmt.Printf("Invalid %s parameter\n", string(config.ContextFrom))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	limit := 0
	if queryParams.Has(string(config.ContextLimit)) {
		limit, err = strconv.Atoi(queryParams.Get(string(config.ContextLimit)))

		if err != nil {
			fmt.Printf("Invalid %s parameter\n", string(config.ContextLimit))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if limit > 50 {
			limit = 50
		}
	}

	// Build the filter
	var filter model.User

	if queryParams.Has("username") {
		filter.Username = queryParams.Get("username")
	}

	// Extract context parameters
	cfg := r.Context().Value(config.ContextConfig).(config.Config)
	abi := r.Context().Value(config.ContextAbi).(string)

	// Select all documents
	users, err := db.SelectUsers(cfg, abi, filter, from, limit)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No users with filter %+v\n", filter)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching users with filter %+v: %s\n", filter, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	if len(users) == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var req model.InsertUserInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Build the new document
	user := model.User{
		Id:       primitive.NewObjectID(),
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		Surname:  req.Surname,
		Birth:    req.Birth,
	}

	// Extract context parameters
	cfg := r.Context().Value(config.ContextConfig).(config.Config)
	abi := r.Context().Value(config.ContextAbi).(string)

	// Insert the new document
	err = db.InsertUser(cfg, abi, user)
	if mongo.IsDuplicateKeyError(err) {
		fmt.Printf("User %+v already exists\n", user)
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err != nil {
		fmt.Printf("Error while inserting user %+v: %s\n", user, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var req model.InsertUserInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract path parameters
	id := r.PathValue(string(config.ContextUserId))
	if id == "" {
		fmt.Printf("Invalid user id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error while converting user id %s: %s\n", id, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Build the new document
	user := model.User{
		Id:       userId,
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		Surname:  req.Surname,
		Birth:    req.Birth,
	}

	// Extract context parameters
	cfg := r.Context().Value(config.ContextConfig).(config.Config)
	abi := r.Context().Value(config.ContextAbi).(string)

	// Update the document
	err = db.UpdateUser(cfg, abi, userId, user)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No users with id %s\n", userId)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Printf("Error while updating user %+v: %s\n", user, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func PatchUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var req model.InsertUserInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract path parameters
	id := r.PathValue(string(config.ContextUserId))
	if id == "" {
		fmt.Printf("Invalid user id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error while converting user id %s: %s\n", id, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(config.ContextConfig).(config.Config)
	abi := r.Context().Value(config.ContextAbi).(string)

	// Select the document
	user, err := db.SelectUser(cfg, abi, userId)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No users with id %s\n", userId)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Printf("Error while searching user with id %s: %s\n", userId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
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

	// Update the document
	err = db.UpdateUser(cfg, abi, userId, user)
	if err != nil {
		fmt.Printf("Error while updating user %+v: %s\n", user, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	id := r.PathValue(string(config.ContextUserId))
	if id == "" {
		fmt.Printf("Invalid user id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error while converting user id %s: %s\n", id, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(config.ContextConfig).(config.Config)
	abi := r.Context().Value(config.ContextAbi).(string)

	// Delete the document
	err = db.DeleteUser(cfg, abi, userId)
	if err != nil {
		fmt.Printf("Error while deleting user with id %s: %s\n", userId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusNoContent)
}
