package handler

import (
	"net/http"

	"mainframe/security/api"
)

func apiKeyByIdHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check of the method request
		switch r.Method {
		case "GET":
			api.GetApiKey(w, r)
		default:
			http.Error(w, r.Method, http.StatusMethodNotAllowed)
		}
	})
}
