package server

import (
	"backend/internal/config"
	openapiV1 "backend/internal/handlers/openapi/v1"
	"backend/internal/logger"
	"fmt"
	oapi "github.com/PostgresContest/openapi/gen/v1"
	"net/http"
)

type Server struct {
	server *oapi.Server
}

func NewProvider(handler *openapiV1.Handler) (*Server, error) {
	server, err := oapi.NewServer(handler)

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
