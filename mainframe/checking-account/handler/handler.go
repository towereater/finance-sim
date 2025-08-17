package handler

import (
	"fmt"
	"net/http"

	"mainframe/checking-account/config"
	mw "mainframe/checking-account/middleware"
)

func SetupRoutes(cfg config.Config, mux *http.ServeMux) {
	// Home path handler
	mux.Handle("/",
		mw.LoggerMiddleware(homeHandler(), cfg))

	// Checking accounts handler
	mux.Handle("/checking-accounts",
		mw.AuthorizedLoggerMiddleware(checkingAccountsHandler(), cfg))
	mux.Handle(fmt.Sprintf("/checking-accounts/{%s}",
		config.ContextAccountId),
		mw.AuthorizedLoggerMiddleware(checkingAccountsByIdHandler(), cfg))

	// Transfers handler
	mux.Handle("/payments/",
		mw.AuthorizedLoggerMiddleware(paymentsHandler(), cfg))
	mux.Handle(fmt.Sprintf("/payments/{%s}",
		config.ContextPaymentId),
		mw.AuthorizedLoggerMiddleware(paymentsByIdHandler(), cfg))
}

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Response output
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello from checking accounts API")
	})
}
