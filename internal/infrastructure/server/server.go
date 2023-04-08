package server

import (
	"fmt"
	"net/http"
	"time"

	"backend/internal/infrastructure/config"
	openapiV1 "backend/internal/infrastructure/handlers/openapi/v1"
	"backend/internal/logger"
)

type SrvProvider interface {
	http.Handler
	BaseRoute() string
}

const (
	readTimeoutSeconds  = 10
	writeTimeoutSeconds = 30
)

func makeServer(providers ...SrvProvider) http.Handler {
	mux := http.NewServeMux()

	for _, provider := range providers {
		root := provider.BaseRoute() + "/"
		mux.Handle(root, http.StripPrefix(provider.BaseRoute(), provider))
	}

	return mux
}

func Invoke(log *logger.Logger, cfg *config.Config, srv *openapiV1.Server) error {
	l := log.WithField("module", "server")

	mux := makeServer(srv)

	listenFullAddr := fmt.Sprintf("%s:%d", cfg.HTTP.Addr, cfg.HTTP.Port)
	HTTPSrv := &http.Server{
		Addr:         listenFullAddr,
		ReadTimeout:  readTimeoutSeconds * time.Second,
		WriteTimeout: writeTimeoutSeconds * time.Second,
		Handler:      mux,
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
