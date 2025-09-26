package handler

import (
	"fmt"
	"net/http"

	"bff/config"
	mw "mainframe-lib/common/middleware"
)

func SetupRoutes(cfg config.Config, mux *http.ServeMux) {
	// Home path handler
	mux.Handle("/",
		mw.LoggerMiddleware(homeHandler(), cfg))

	// Users handler
	mux.Handle("/user/register",
		mw.AuthorizedLoggerMiddleware(userRegisterHandler(), cfg, securityAuth()))
	mux.Handle("/user/login",
		mw.AuthorizedLoggerMiddleware(userLoginHandler(), cfg, securityAuth()))
	mux.Handle("/user/password",
		mw.AuthorizedLoggerMiddleware(userPasswordHandler(), cfg, jwtAuth()))

	// Accounts handler
	mux.Handle("/accounts",
		mw.AuthorizedLoggerMiddleware(accountsHandler(), cfg, jwtAuth()))

	// Payments handler
	mux.Handle("/payments",
		mw.AuthorizedLoggerMiddleware(paymentsHandler(), cfg, jwtAuth()))
	mux.Handle(fmt.Sprintf("/payments/{%s}",
		config.ContextPaymentId),
		mw.AuthorizedLoggerMiddleware(paymentByIdHandler(), cfg, jwtAuth()))

	// Stocks handler
	mux.Handle("/stocks",
		mw.AuthorizedLoggerMiddleware(stocksHandler(), cfg, jwtAuth()))
	mux.Handle(fmt.Sprintf("/stocks/{%s}",
		config.ContextIsin),
		mw.AuthorizedLoggerMiddleware(stockByIsinHandler(), cfg, jwtAuth()))

	// Orders handler
	mux.Handle("/orders",
		mw.AuthorizedLoggerMiddleware(ordersHandler(), cfg, jwtAuth()))
	mux.Handle(fmt.Sprintf("/orders/{%s}",
		config.ContextOrderId),
		mw.AuthorizedLoggerMiddleware(orderByIdHandler(), cfg, jwtAuth()))
}

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Response output
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello from bff API")
	})
}
