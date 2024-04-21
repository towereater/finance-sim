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
		//TODO: INSERT A NEW ACCOUNT IN THE OWNER ACCOUNTS LIST
		api.InsertAccount(w, r)
	default:
		http.Error(w, r.Method, http.StatusMethodNotAllowed)
	}
}

// Handles users API functions
func accountsByIdHandler(w http.ResponseWriter, r *http.Request) {
	// Check of the method request
	switch r.Method {
	case "GET":
		api.GetAccount(w, r)
	case "PUT":
		//TODO: UPDATE THE ACCOUNT OF THE OWNER ACCOUNTS LIST
		api.UpdateAccount(w, r)
	case "PATCH":
		//TODO: UPDATE THE ACCOUNT OF THE OWNER ACCOUNTS LIST IF OWNER IS SET
		api.PatchAccount(w, r)
	case "DELETE":
		//TODO: REMOVE THE ACCOUNT FROM THE OWNER ACCOUNTS LIST
		api.DeleteAccount(w, r)
	default:
		http.Error(w, r.Method, http.StatusMethodNotAllowed)
	}
}
