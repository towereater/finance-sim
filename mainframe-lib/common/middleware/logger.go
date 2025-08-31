package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func logger() Adapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Log the request
			fmt.Printf("%s %s request at %s\n",
				time.Now().UTC().Format("2006-01-02T15:04:05"),
				r.Method,
				r.URL,
			)

			next.ServeHTTP(w, r)
		})
	}
}
