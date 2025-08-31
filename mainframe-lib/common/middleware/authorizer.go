package middleware

import (
	"context"
	"fmt"
	"mainframe-lib/common/config"
	"net/http"
	"strings"
)

func authorizer() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract header values
			auth := r.Header.Get("Authorization")
			components := strings.Split(auth, ":")

			if len(components) < 2 {
				fmt.Printf("Invalid authorization %s\n", auth)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Add a basic configuration to the request context
			ctx := context.WithValue(r.Context(), config.ContextAbi, components[0])
			ctx = context.WithValue(ctx, config.ContextApiKey, components[1])
			ctx = context.WithValue(ctx, config.ContextAuth, auth)

			newReq := r.WithContext(ctx)
			h.ServeHTTP(w, newReq)
		})
	}
}
