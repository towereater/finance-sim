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

	// Dossiers handler
	mux.Handle("/dossiers",
		mw.AuthorizedLoggerMiddleware(dossiersHandler(), cfg, securityAuth()))
	mux.Handle(fmt.Sprintf("/dossiers/{%s}",
		config.ContextDossier),
		mw.AuthorizedLoggerMiddleware(dossiersByIdHandler(), cfg, securityAuth()))
}

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Response output
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello from dossiers API")
	})
}
