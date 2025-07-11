package handler

import (
	"fmt"
	"net/http"

	"mainframe/user/config"
	mw "mainframe/user/middleware"
)

func SetupRoutes(cfg config.Config, mux *http.ServeMux) {
	// Home path handler
	mux.Handle("/",
		mw.LoggerMiddleware(homeHandler(), cfg))

	// Users handler
	mux.Handle("/users",
		mw.AuthorizedLoggerMiddleware(usersHandler(), cfg))
	mux.Handle(fmt.Sprintf("/users/{%s}",
		config.ContextUserId),
		mw.AuthorizedLoggerMiddleware(usersByIdHandler(), cfg))
	mux.Handle(fmt.Sprintf("/users/{%s}/accounts/{%s}",
		config.ContextUserId,
		config.ContextAccountId),
		mw.AuthorizedLoggerMiddleware(userAccountsByIdHandler(), cfg))
}

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Response output
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello from users API")
	})
}
