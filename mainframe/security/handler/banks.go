package handler

import (
	"mainframe/security/api"
	"net/http"
)

func banksHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check of the method request
		switch r.Method {
		case "POST":
			api.InsertBank(w, r)
		default:
			http.Error(w, r.Method, http.StatusMethodNotAllowed)
		}
	})
}

func bankByIdHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check of the method request
		switch r.Method {
		case "GET":
			api.GetBank(w, r)
		case "DELETE":
			api.DeleteBank(w, r)
		default:
			http.Error(w, r.Method, http.StatusMethodNotAllowed)
		}
	})
}
