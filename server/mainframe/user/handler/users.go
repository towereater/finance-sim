package handler

import (
	"fmt"
	"net/http"

	"mainframe/user/api"
)

func HandleRequests() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/{userId}", userByIdHandler)
	http.HandleFunc("/users/{userId}/add-account", userAddAccountHandler)
	http.HandleFunc("/users/{userId}/remove-account", userRemoveAccountHandler)
}

// Handles home path
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Response output
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello from users API")
}

// Handles users API functions
func usersHandler(w http.ResponseWriter, r *http.Request) {
	// Check of the method request
	switch r.Method {
	case "GET":
		api.GetUsers(w, r)
	case "POST":
		api.InsertUser(w, r)
	default:
		http.Error(w, r.Method, http.StatusMethodNotAllowed)
	}
}

// Handles users API functions
func userByIdHandler(w http.ResponseWriter, r *http.Request) {
	var handledPath = "/users/{id}"

	// Check of the method request
	switch r.Method {
	case "GET":
		api.GetUser(w, r, handledPath)
	case "PUT":
		api.UpdateUser(w, r, handledPath)
	case "PATCH":
		api.PatchUser(w, r, handledPath)
	case "DELETE":
		api.DeleteUser(w, r, handledPath)
	default:
		http.Error(w, r.Method, http.StatusMethodNotAllowed)
	}
}

// Handles user accounts API functions
func userAddAccountHandler(w http.ResponseWriter, r *http.Request) {
	var handledPath = "/users/{id}/add-account"

	// Check of the method request
	switch r.Method {
	case "PATCH":
		api.AddAccount(w, r, handledPath)
	default:
		http.Error(w, r.Method, http.StatusMethodNotAllowed)
	}
}

// Handles user accounts API functions
func userRemoveAccountHandler(w http.ResponseWriter, r *http.Request) {
	var handledPath = "/users/{id}/remove-account"

	// Check of the method request
	switch r.Method {
	case "PATCH":
		api.RemoveAccount(w, r, handledPath)
	default:
		http.Error(w, r.Method, http.StatusMethodNotAllowed)
	}
}
