package middleware

import (
	"net/http"
)

type Adapter func(http.Handler) http.Handler

func chainMiddleware(h http.Handler, adapters ...Adapter) http.Handler {
	for i := len(adapters) - 1; i >= 0; i-- {
		h = adapters[i](h)
	}

	return h
}

func LoggerMiddleware(h http.Handler, cfg any) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		adapters := []Adapter{
			logger(),
			addConfig(cfg),
			addType(),
		}
		chainMiddleware(h, adapters...).ServeHTTP(w, r)
	})
}

func AuthorizedLoggerMiddleware(h http.Handler, cfg any) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		adapters := []Adapter{
			logger(),
			addConfig(cfg),
			authorizer(),
			addType(),
		}
		chainMiddleware(h, adapters...).ServeHTTP(w, r)
	})
}
