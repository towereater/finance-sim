package handler

import (
	"fmt"
	"net/http"

	mw "mainframe-lib/common/middleware"
	"mainframe/checking-account/config"
)

func SetupRoutes(cfg config.Config, mux *http.ServeMux) {
	// Home path handler
	mux.Handle("/",
		mw.LoggerMiddleware(homeHandler(), cfg))

	// Checking accounts handler
	mux.Handle("/checking-accounts",
		mw.AuthorizedLoggerMiddleware(checkingAccountsHandler(), cfg, securityAuth()))
	mux.Handle(fmt.Sprintf("/checking-accounts/{%s}",
		config.ContextAccountId),
		mw.AuthorizedLoggerMiddleware(checkingAccountsByIdHandler(), cfg, securityAuth()))

	// Transfers handler
	mux.Handle("/payments",
		mw.AuthorizedLoggerMiddleware(paymentsHandler(), cfg, securityAuth()))
	mux.Handle(fmt.Sprintf("/payments/{%s}",
		config.ContextPaymentId),
		mw.AuthorizedLoggerMiddleware(paymentsByIdHandler(), cfg, securityAuth()))
}

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Response output
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello from checking accounts API")
	})
}
