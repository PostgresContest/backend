package middlewares

import (
	"backend/internal/types"

	"github.com/ogen-go/ogen/middleware"
)

func RecoverMiddleware(log types.Logger) middleware.Middleware {
	return func(req middleware.Request, next middleware.Next) (middleware.Response, error) {
		defer func() {
			if r := recover(); r != nil {
				log.Warn("recovered", r)
			}
		}()

		return next(req)
	}
}
