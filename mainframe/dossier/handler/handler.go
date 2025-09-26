package handler

import (
	"fmt"
	"net/http"

	mw "mainframe-lib/common/middleware"
	"mainframe/dossier/config"
)

func SetupRoutes(cfg config.Config, mux *http.ServeMux) {
	// Home path handler
	mux.Handle("/",
		mw.LoggerMiddleware(homeHandler(), cfg))

	// Stocks handler
	mux.Handle("/stocks",
		mw.AuthorizedLoggerMiddleware(stocksHandler(), cfg, securityAuth()))
	mux.Handle(fmt.Sprintf("/stocks/{%s}",
		config.ContextIsin),
		mw.AuthorizedLoggerMiddleware(stockByIsinHandler(), cfg, securityAuth()))

	// Dossiers handler
	mux.Handle("/dossiers",
		mw.AuthorizedLoggerMiddleware(dossiersHandler(), cfg, securityAuth()))
	mux.Handle(fmt.Sprintf("/dossiers/{%s}",
		config.ContextDossierId),
		mw.AuthorizedLoggerMiddleware(dossierByIdHandler(), cfg, securityAuth()))

	// Orders handler
	mux.Handle("/orders",
		mw.AuthorizedLoggerMiddleware(ordersHandler(), cfg, securityAuth()))
	mux.Handle(fmt.Sprintf("/orders/{%s}",
		config.ContextOrderId),
		mw.AuthorizedLoggerMiddleware(orderByIdHandler(), cfg, securityAuth()))
}

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Response output
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello from dossiers API")
	})
}
