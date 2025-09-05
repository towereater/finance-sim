package handler

import (
	"context"
	"fmt"
	com "mainframe-lib/common/config"
	mw "mainframe-lib/common/middleware"
	sec "mainframe-lib/security/service"
	"mainframe/checking-account/config"
	"net/http"
	"strings"
)

func securityAuth() mw.Adapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract header values
			auth := r.Header.Get("Authorization")
			components := strings.Split(auth, ":")

			if len(components) < 2 {
				fmt.Printf("Invalid authorization %s\n", auth)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Extract context parameters
			cfg := r.Context().Value(com.ContextConfig).(config.Config)

			// Extract api key
			apiKey := components[1]

			// Check api key existence
			_, err := sec.GetUserByApiKey(
				cfg.Services.Security,
				cfg.Services.Timeout,
				auth,
				apiKey,
			)
			if err != nil {
				fmt.Printf("Error while validating api key %s: %s\n", apiKey, err.Error())
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), com.ContextAbi, components[0])
			ctx = context.WithValue(ctx, com.ContextApiKey, components[1])
			ctx = context.WithValue(ctx, com.ContextAuth, auth)

			newReq := r.WithContext(ctx)
			next.ServeHTTP(w, newReq)
		})
	}
}
