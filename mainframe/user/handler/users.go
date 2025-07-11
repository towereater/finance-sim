package handler

import (
	"net/http"

	"mainframe/user/api"
)

func usersHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check of the method request
		switch r.Method {
		case "GET":
			api.GetUsers(w, r)
		case "POST":
			api.InsertUser(w, r)
		default:
			http.Error(w, r.Method, http.StatusMethodNotAllowed)
		}
	})
}

func usersByIdHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	})
}

func userAccountsByIdHandler() http.Handler {
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
