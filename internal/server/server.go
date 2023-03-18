package server

import (
	"backend/internal/config"
	"backend/internal/logger"
	"fmt"
	oapi "github.com/PostgresContest/openapi/gen/v1"
	"net/http"
)

type Server struct {
	server *oapi.Server
}

func NewProvider() (*Server, error) {
	server, err := oapi.NewServer(oapi.UnimplementedHandler{})

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
