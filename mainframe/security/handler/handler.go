package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	com "mainframe-lib/common/config"
	mw "mainframe-lib/common/middleware"
	"mainframe/security/config"
)

func SetupRoutes(cfg config.Config, mux *http.ServeMux) {
	// Home path handler
	mux.Handle("/",
		mw.LoggerMiddleware(homeHandler(), cfg))

	// Users handler
	mux.Handle("/users",
		mw.AuthorizedLoggerMiddleware(usersHandler(), cfg, basicAuth()))
	mux.Handle(fmt.Sprintf("/users/{%s}",
		config.ContextUserId),
		mw.AuthorizedLoggerMiddleware(userByIdHandler(), cfg, basicAuth()))

	// Api keys handler
	mux.Handle(fmt.Sprintf("/api-keys/{%s}",
		config.ContextApiKey),
		mw.AuthorizedLoggerMiddleware(apiKeyByIdHandler(), cfg, basicAuth()))
}

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Response output
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello from security API")
	})
}

func basicAuth() mw.Adapter {
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

			// Add authorization headers to the request context
			ctx := context.WithValue(r.Context(), com.ContextAbi, components[0])
			ctx = context.WithValue(ctx, com.ContextApiKey, components[1])
			ctx = context.WithValue(ctx, com.ContextAuth, auth)

			newReq := r.WithContext(ctx)
			next.ServeHTTP(w, newReq)
		})
	}
}
