package middleware

import (
	"context"
	"mainframe/security/config"
	"net/http"
)

func authorizer() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract header values
			auth := r.Header.Get("Authorization")

			// Add a basic configuration to the request context
			ctx := context.WithValue(r.Context(), config.ContextAbi, "09999")
			ctx = context.WithValue(ctx, config.ContextCab, "00099")
			ctx = context.WithValue(ctx, config.ContextAuth, auth)

			newReq := r.WithContext(ctx)
			h.ServeHTTP(w, newReq)
		})
	}
}
