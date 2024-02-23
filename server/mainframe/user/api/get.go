package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"mainframe/user/db"
	"mainframe/user/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get user API function
func GetUser(w http.ResponseWriter, r *http.Request, urlModel string) {
	// Extraction of extra parameters
	pathParams := getPathParams(r.URL, urlModel)

	id, err := primitive.ObjectIDFromHex(pathParams["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	// Execution of the request
	user, err := db.SelectUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response output
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Get users API function
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Extraction of extra parameters
	queryParams := r.URL.Query()

	fromId := primitive.NilObjectID
	if queryParams.Has("fromId") {
		fromId, _ = primitive.ObjectIDFromHex(queryParams.Get("fromId"))
	}

	limit := 0
	if queryParams.Has("limit") {
		limit, _ = strconv.Atoi(queryParams.Get("limit"))
	}

	order := 1
	if queryParams.Has("order") {
		order, _ = strconv.Atoi(queryParams.Get("order"))
	}

	// Building the filter
	var filter model.User

	if queryParams.Has("username") {
		filter.Name = queryParams.Get("username")
	}

	if queryParams.Has("password") {
		filter.Name = queryParams.Get("password")
	}

	if queryParams.Has("name") {
		filter.Name = queryParams.Get("name")
	}

	if queryParams.Has("surname") {
		filter.Surname = queryParams.Get("surname")
	}

	// Execution of the request
	user, err := db.SelectUsers(filter, fromId, limit, order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Response output
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
