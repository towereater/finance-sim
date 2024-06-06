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
	http.HandleFunc("/users/{userId}/accounts/{accountId}", userAccountsByIdHandler)
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
	// Check of the method request
	switch r.Method {
	case "GET":
		api.GetUser(w, r)
	case "PUT":
		api.UpdateUser(w, r)
	case "PATCH":
		api.PatchUser(w, r)
	case "DELETE":
		api.DeleteUser(w, r)
	default:
		http.Error(w, r.Method, http.StatusMethodNotAllowed)
	}
}

// Handles user accounts API functions
func userAccountsByIdHandler(w http.ResponseWriter, r *http.Request) {
	// Check of the method request
	switch r.Method {
	case "POST":
		api.AddAccount(w, r)
	case "DELETE":
		api.RemoveAccount(w, r)
	default:
		http.Error(w, r.Method, http.StatusMethodNotAllowed)
	}
}
