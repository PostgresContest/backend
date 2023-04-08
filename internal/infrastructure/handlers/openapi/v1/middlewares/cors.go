package middlewares

import (
	"net/http"

	"github.com/rs/cors"
)

func isOriginAllowed(origin string, allowedOrigins []string) bool {
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			return true
		}
	}

	return false
}

func CORSMiddleware(isDevMode bool, allowedOrigins []string) func(h http.Handler) http.Handler {
	rsCors := cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool {
			return isDevMode || isOriginAllowed(origin, allowedOrigins)
		},
		AllowedHeaders: []string{"Authorization"},
	})

	return rsCors.Handler
}
