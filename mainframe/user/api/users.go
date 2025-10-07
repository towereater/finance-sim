package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	com "mainframe-lib/common/config"
	usr "mainframe-lib/user/model"
	"mainframe/user/config"
	"mainframe/user/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	userId := r.PathValue(string(config.ContextUserId))
	if userId == "" || len(userId) != 24 {
		fmt.Printf("Invalid user id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)

	// Select the document
	user, err := db.SelectUser(cfg.DBConfig, abi, userId)
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

	from := queryParams.Get(string(config.ContextFrom))
	if from != "" && len(from) != 24 {
		fmt.Printf("Invalid %s parameter\n", string(config.ContextFrom))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	limit := 50
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
	var filter usr.User
	filter.Username = queryParams.Get("username")
	filter.Password = queryParams.Get("password")

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)

	// Select all documents
	users, err := db.SelectUsers(cfg.DBConfig, abi, filter, from, limit)
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
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request
	var req usr.InsertUserInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Username == "" {
		fmt.Printf("Invalid user username\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Password == "" {
		fmt.Printf("Invalid user password\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		fmt.Printf("Invalid user name\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Surname == "" {
		fmt.Printf("Invalid user surname\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Birth == "" || len(req.Birth) != 10 {
		fmt.Printf("Invalid user birth\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Cab == "" || len(req.Cab) != 5 {
		fmt.Printf("Invalid user cab\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)

	// Build the new document
	user := usr.User{
		Id:       primitive.NewObjectID().Hex(),
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		Surname:  req.Surname,
		Birth:    req.Birth,
		Cab:      req.Cab,
	}

	// Insert the new document
	err = db.InsertUser(cfg.DBConfig, abi, user)
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
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func PatchUser(w http.ResponseWriter, r *http.Request) {
	// Extract path parameters
	userId := r.PathValue(string(config.ContextUserId))
	if userId == "" || len(userId) != 24 {
		fmt.Printf("Invalid user id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse the request
	var req usr.UpdateUserInput
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("Could not convert request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)

	// Select the document
	user, err := db.SelectUser(cfg.DBConfig, abi, userId)
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

	// Generate the updated document
	if req.Password != "" {
		user.Password = req.Password
	}

	// Update the document
	err = db.UpdateUser(cfg.DBConfig, abi, userId, user)
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
	userId := r.PathValue(string(config.ContextUserId))
	if userId == "" || len(userId) != 24 {
		fmt.Printf("Invalid user id value\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Extract context parameters
	cfg := r.Context().Value(com.ContextConfig).(config.Config)
	abi := r.Context().Value(com.ContextAbi).(string)

	// Delete the document
	err := db.DeleteUser(cfg.DBConfig, abi, userId)
	if err != nil {
		fmt.Printf("Error while deleting user with id %s: %s\n", userId, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusNoContent)
}
