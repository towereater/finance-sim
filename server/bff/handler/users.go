package handler

import (
	"fmt"
	"net/http"

	"bff/api/user"
)

// Handles all the API functions for this service
func HandleRequests() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/users/register", usersRegisterHandler)
	http.HandleFunc("/users/login", usersLoginHandler)
	http.HandleFunc("/users/change-password", usersChangePasswordHandler)
	http.HandleFunc("/users/delete-user", usersDeleteUserHandler)
}

// Handles home path
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Response output
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello from users BFF")
}

// Handles users API functions
func usersRegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Check of the method request
	switch r.Method {
	case "POST":
		user.RegisterUser(w, r)
	default:
		http.Error(w, r.Method, http.StatusMethodNotAllowed)
	}
}

// Handles users API functions
func usersLoginHandler(w http.ResponseWriter, r *http.Request) {
	// Check of the method request
	switch r.Method {
	case "POST":
		user.LoginUser(w, r)
	default:
		http.Error(w, r.Method, http.StatusMethodNotAllowed)
	}
}

// Handles users API functions
func usersChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	// Check of the method request
	switch r.Method {
	case "POST":
		user.ChangePasswordUser(w, r)
	default:
		http.Error(w, r.Method, http.StatusMethodNotAllowed)
	}
}

// Handles users API functions
func usersDeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Check of the method request
	switch r.Method {
	case "POST":
		user.DeleteUser(w, r)
	default:
		http.Error(w, r.Method, http.StatusMethodNotAllowed)
	}
}
