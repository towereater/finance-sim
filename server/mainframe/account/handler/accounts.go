package handler

import (
	"fmt"
	"net/http"

	"mainframe/account/api"
)

// Handles all the API functions for this service
func HandleRequests() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/accounts", accountsHandler)
	http.HandleFunc("/accounts/{accountId}", accountsByIdHandler)
}

// Handles home path
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Response output
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello from accounts API")
}

// Handles users API functions
func accountsHandler(w http.ResponseWriter, r *http.Request) {
	// Check of the method request
	switch r.Method {
	case "GET":
		api.GetAccounts(w, r)
	case "POST":
		api.InsertAccount(w, r)
	default:
		http.Error(w, r.Method, http.StatusMethodNotAllowed)
	}
}

// Handles users API functions
func accountsByIdHandler(w http.ResponseWriter, r *http.Request) {
	var handledPath = "/accounts/{id}"

	// Check of the method request
	switch r.Method {
	case "GET":
		api.GetAccount(w, r, handledPath)
	case "PUT":
		api.UpdateAccount(w, r, handledPath)
	case "PATCH":
		api.PatchAccount(w, r, handledPath)
	case "DELETE":
		api.DeleteAccount(w, r, handledPath)
	default:
		http.Error(w, r.Method, http.StatusMethodNotAllowed)
	}
}
