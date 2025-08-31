package middleware

import (
	"context"
	"mainframe-lib/common/config"
	"net/http"
)

func addConfig(cfg any) Adapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Add a basic configuration to the request context
			ctx := context.WithValue(r.Context(), config.ContextConfig, cfg)
			newReq := r.WithContext(ctx)

			next.ServeHTTP(w, newReq)
		})
	}
}
