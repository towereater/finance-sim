package handler

import (
	"mainframe/security/api"
	"net/http"
)

func usersHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check of the method request
		switch r.Method {
		case "POST":
			api.InsertUser(w, r)
		default:
			http.Error(w, r.Method, http.StatusMethodNotAllowed)
		}
	})
}

func userByIdHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check of the method request
		switch r.Method {
		case "DELETE":
			api.DeleteUser(w, r)
		default:
			http.Error(w, r.Method, http.StatusMethodNotAllowed)
		}
	})
}
