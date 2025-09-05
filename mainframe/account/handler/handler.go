package handler

import (
	"fmt"
	"net/http"

	mw "mainframe-lib/common/middleware"
	"mainframe/account/config"
)

func SetupRoutes(cfg config.Config, mux *http.ServeMux) {
	// Home path handler
	mux.Handle("/",
		mw.LoggerMiddleware(homeHandler(), cfg))

	// Accounts handler
	mux.Handle("/accounts",
		mw.AuthorizedLoggerMiddleware(accountsHandler(), cfg, securityAuth()))
	mux.Handle(fmt.Sprintf("/accounts/services/{%s}",
		config.ContextService),
		mw.AuthorizedLoggerMiddleware(accountsByServiceHandler(), cfg, securityAuth()))
	mux.Handle(fmt.Sprintf("/accounts/services/{%s}/accounts/{%s}",
		config.ContextService,
		config.ContextAccount),
		mw.AuthorizedLoggerMiddleware(accountsByIdHandler(), cfg, securityAuth()))
}

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Response output
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello from accounts API")
	})
}
