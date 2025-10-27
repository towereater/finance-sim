package handler

import (
	"net/http"

	"mainframe/checking-account/api"
)

func checkingAccountsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check of the method request
		switch r.Method {
		case "GET":
			api.GetCheckingAccounts(w, r)
		case "POST":
			api.InsertCheckingAccount(w, r)
		default:
			http.Error(w, r.Method, http.StatusMethodNotAllowed)
		}
	})
}

func checkingAccountsByIdHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check of the method request
		switch r.Method {
		case "GET":
			api.GetCheckingAccount(w, r)
		case "DELETE":
			api.DeleteCheckingAccount(w, r)
		default:
			http.Error(w, r.Method, http.StatusMethodNotAllowed)
		}
	})
}

func checkingAccountsValueHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check of the method request
		switch r.Method {
		case "POST":
			api.AddCheckingValue(w, r)
		default:
			http.Error(w, r.Method, http.StatusMethodNotAllowed)
		}
	})
}
