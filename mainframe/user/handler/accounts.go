package handler

import (
	"mainframe/user/api"
	"net/http"
)

func userAccountsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check of the method request
		switch r.Method {
		case "POST":
			api.AddAccount(w, r)
		case "DELETE":
			api.RemoveAccount(w, r)
		default:
			http.Error(w, r.Method, http.StatusMethodNotAllowed)
		}
	})
}
