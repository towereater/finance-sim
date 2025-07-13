package middleware

import (
	"context"
	"mainframe/account/config"
	"net/http"
)

func authorizer() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Add a basic configuration to the request context
			ctx := context.WithValue(r.Context(), config.ContextAbi, "06270")
			newReq := r.WithContext(ctx)
			h.ServeHTTP(w, newReq)
		})
	}
}
