package middleware

import (
	"net/http"
)

func addType() Adapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Add content type header to the response
			w.Header().Set("Content-Type", "application/json")

			next.ServeHTTP(w, r)
		})
	}
}
