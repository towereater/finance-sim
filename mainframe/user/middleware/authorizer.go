package middleware

import (
	"context"
	"fmt"
	"mainframe/user/config"
	"mainframe/user/service"
	"net/http"
)

func authorizer() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract header values
			auth := r.Header.Get("Authorization")

			// Extract context parameters
			cfg := r.Context().Value(config.ContextConfig).(config.Config)

			// Check api key existence
			apiKey, err := service.GetApiKey(cfg, auth)
			if err != nil {
				fmt.Printf("Error while api key %s: %s\n", auth, err.Error())
				w.WriteHeader(http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), config.ContextAbi, apiKey.Abi)
			ctx = context.WithValue(ctx, config.ContextCab, apiKey.Cab)
			ctx = context.WithValue(ctx, config.ContextAuth, apiKey.Id)

			newReq := r.WithContext(ctx)
			h.ServeHTTP(w, newReq)
		})
	}
}
