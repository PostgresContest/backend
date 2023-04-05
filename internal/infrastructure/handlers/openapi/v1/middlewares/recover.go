package middlewares

import (
	"backend/internal/types"
	"net/http"
	"runtime/debug"
)

func getStack() []byte {
	return debug.Stack()
}

func RecoverMiddleware(log types.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				defer func() {
					if r := recover(); r != nil {
						log.Errorf("recovered; caused: %v; stacktrace: %s", r, getStack())
					}
				}()

				h.ServeHTTP(w, r)
			},
		)
	}
}
