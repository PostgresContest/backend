package server

import (
	"backend/internal/infrastructure/config"
	openapiV1 "backend/internal/infrastructure/handlers/openapi/v1"
	"backend/internal/middlewares"
	"fmt"
	"net/http"
	"time"

	"backend/internal/infrastructure/auth"
	"backend/internal/logger"
	oapi "github.com/PostgresContest/openapi/gen/v1"
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

const (
	readTimeoutSeconds  = 10
	writeTimeoutSeconds = 30
)

func Invoke(log *logger.Logger, cfg *config.Config, srv *Server) error {
	l := log.WithField("module", "server")
	listenFullAddr := fmt.Sprintf("%s:%d", cfg.HTTP.Addr, cfg.HTTP.Port)

	HTTPSrv := &http.Server{
		Addr:         listenFullAddr,
		ReadTimeout:  readTimeoutSeconds * time.Second,
		WriteTimeout: writeTimeoutSeconds * time.Second,
		Handler:      srv.server,
	}

	l.Info("starting server")

	go func() {
		err := HTTPSrv.ListenAndServe()
		if err != nil {
			panic(fmt.Sprintf("cannot start server: %e", err))
		}
	}()

	return nil
}
