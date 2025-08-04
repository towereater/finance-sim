package handler

import (
	"fmt"
	"net/http"

	"mainframe/dossier/config"
	mw "mainframe/dossier/middleware"
)

func SetupRoutes(cfg config.Config, mux *http.ServeMux) {
	// Home path handler
	mux.Handle("/",
		mw.LoggerMiddleware(homeHandler(), cfg))

	// Dossiers handler
	mux.Handle("/dossiers",
		mw.AuthorizedLoggerMiddleware(dossiersHandler(), cfg))
	mux.Handle(fmt.Sprintf("/dossiers/{%s}",
		config.ContextDossier),
		mw.AuthorizedLoggerMiddleware(dossiersByIdHandler(), cfg))
}

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Response output
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello from dossiers API")
	})
}
