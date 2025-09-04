package handler

import (
	"fmt"
	"net/http"

	mw "mainframe-lib/common/middleware"
	"mainframe/user/config"
)

func SetupRoutes(cfg config.Config, mux *http.ServeMux) {
	// Home path handler
	mux.Handle("/",
		mw.LoggerMiddleware(homeHandler(), cfg))

	// Users handler
	mux.Handle("/users",
		mw.AuthorizedLoggerMiddleware(usersHandler(), cfg, securityAuth()))
	mux.Handle(fmt.Sprintf("/users/{%s}",
		config.ContextUserId),
		mw.AuthorizedLoggerMiddleware(userByIdHandler(), cfg, securityAuth()))
	mux.Handle(fmt.Sprintf("/users/{%s}/accounts",
		config.ContextUserId),
		mw.AuthorizedLoggerMiddleware(userAccountsHandler(), cfg, securityAuth()))
}

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Response output
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello from users API")
	})
}
