package v1

import (
	"net/http"

	"backend/internal/infrastructure/auth"
	"backend/internal/infrastructure/config"
	"backend/internal/infrastructure/handlers/openapi/v1/middlewares"
	"backend/internal/logger"
	oapi "github.com/PostgresContest/openapi/go/gen/v1"
)

type Server struct {
	http.Handler
}

func (s *Server) BaseRoute() string {
	return "/api"
}

func NewServerProvider(
	log *logger.Logger,
	handler *Handler,
	security *auth.Security,
	cfg *config.Config,
) (*Server, error) {
	srv, err := oapi.NewServer(
		handler,
		security,
	)
	if err != nil {
		return nil, err
	}

	recoverMiddleware := middlewares.RecoverMiddleware(log.WithField("module", "middleware.recover"))
	corsMiddleware := middlewares.CORSMiddleware(cfg.IsDevMode(), cfg.CORS.AllowedOrigins)

	return &Server{
		recoverMiddleware(
			corsMiddleware(srv),
		),
	}, err
}
