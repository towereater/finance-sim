package handler

import (
	"fmt"
	"net/http"

	"mainframe/account/config"
	mw "mainframe/account/middleware"
)

func SetupRoutes(cfg config.Config, mux *http.ServeMux) {
	// Home path handler
	mux.Handle("/",
		mw.LoggerMiddleware(homeHandler(), cfg))

	// Users handler
	mux.Handle("/accounts",
		mw.AuthorizedLoggerMiddleware(accountsHandler(), cfg))
	mux.Handle(fmt.Sprintf("/accounts/{%s}",
		config.ContextAccountId),
		mw.AuthorizedLoggerMiddleware(accountsByIdHandler(), cfg))
}

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Response output
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello from accounts API")
	})
}
