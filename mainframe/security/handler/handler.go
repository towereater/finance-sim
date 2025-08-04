package handler

import (
	"fmt"
	"net/http"

	"mainframe/security/config"
	mw "mainframe/security/middleware"
)

func SetupRoutes(cfg config.Config, mux *http.ServeMux) {
	// Home path handler
	mux.Handle("/",
		mw.LoggerMiddleware(homeHandler(), cfg))

	// Api keys handler
	mux.Handle("/api-keys",
		mw.AuthorizedLoggerMiddleware(apiKeysHandler(), cfg))
	mux.Handle(fmt.Sprintf("/api-keys/{%s}",
		config.ContextApiKey),
		mw.AuthorizedLoggerMiddleware(apiKeysByIdHandler(), cfg))
}

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Response output
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello from security API")
	})
}
