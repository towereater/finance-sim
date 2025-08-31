package middleware

import (
	"net/http"
)

// Middleware main type
type Adapter func(http.Handler) http.Handler

func chainMiddleware(h http.Handler, adapters ...Adapter) http.Handler {
	for i := len(adapters) - 1; i >= 0; i-- {
		h = adapters[i](h)
	}

	return h
}

func LoggerMiddleware(h http.Handler, cfg any) http.Handler {
	adapters := []Adapter{
		logger(),
		addConfig(cfg),
		addType(),
	}

	return chainMiddleware(h, adapters...)
}

func AuthorizedLoggerMiddleware(h http.Handler, cfg any, authorizer Adapter) http.Handler {
	adapters := []Adapter{
		logger(),
		addConfig(cfg),
		authorizer,
		addType(),
	}

	return chainMiddleware(h, adapters...)
}
