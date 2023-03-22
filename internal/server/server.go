package server

import (
	"backend/internal/config"
	openapiV1 "backend/internal/handlers/openapi/v1"
	"backend/internal/infrastructure/auth"
	"backend/internal/logger"
	"backend/internal/middlewares"
	"fmt"
	oapi "github.com/PostgresContest/openapi/gen/v1"
	"net/http"
)

type Server struct {
	server *oapi.Server
}

func NewProvider(
	log *logger.Logger,
	handler *openapiV1.Handler,
	security *auth.Security,
) (*Server, error) {
	server, err := oapi.NewServer(
		handler,
		security,
		oapi.WithMiddleware(
			middlewares.RecoverMiddleware(log.WithField("module", "middleware.recover")),
		),
	)

	if err != nil {
		return nil, err
	}

	return &Server{server: server}, nil
}

func Invoke(log *logger.Logger, cfg *config.Config, srv *Server) error {
	l := log.WithField("module", "server")
	listenFullAddr := fmt.Sprintf("%s:%d", cfg.Http.Addr, cfg.Http.Port)

	l.Info("starting server")
	err := http.ListenAndServe(listenFullAddr, srv.server)
	if err != nil {
		l.Fatalf("cannot start server: %e", err)
	}
	return err
}
