package app

import (
	"backend/internal/config"
	openapiV1 "backend/internal/handlers/openapi/v1"
	"backend/internal/handlers/openapi/v1/modules/auth"
	dbPrivate "backend/internal/infrastructure/db/private"

	"backend/internal/logger"
	userRepository "backend/internal/repositories/user"
	"backend/internal/server"
	"go.uber.org/fx"
)

func getProvidersAndInvokers() ([]any, []any) {
	providers := []any{
		config.NewProvider,
		logger.NewProvider,

		dbPrivate.NewProvider,

		userRepository.NewProvider,

		auth.NewProvider,

		openapiV1.NewProvider,
		server.NewProvider,
	}
	invokers := []any{
		server.Invoke,
	}

	return providers, invokers
}

func getFx() *fx.App {
	providers, invokers := getProvidersAndInvokers()
	f := fx.New(
		fx.Provide(providers...),
		fx.Invoke(invokers...),
	)
	return f
}
