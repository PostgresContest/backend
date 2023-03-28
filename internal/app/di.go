package app

import (
	"backend/internal/config"
	openapiV1 "backend/internal/handlers/openapi/v1"
	authModule "backend/internal/handlers/openapi/v1/modules/auth"
	userModule "backend/internal/handlers/openapi/v1/modules/user"
	"backend/internal/infrastructure/auth"
	dbPrivate "backend/internal/infrastructure/db/private"
	dbUserAccess "backend/internal/infrastructure/db/user_access"
	userRepository "backend/internal/infrastructure/repositories/user"
	"backend/internal/logger"
	"backend/internal/server"

	"go.uber.org/fx"
)

func getProvidersAndInvokers() ([]any, []any) {
	providers := []any{
		config.NewProvider,
		logger.NewProvider,

		dbPrivate.NewProvider,
		dbUserAccess.NewProvider,
		auth.NewAccessTokenProvider,
		auth.NewSecurityProvider,

		userRepository.NewProvider,

		authModule.NewProvider,
		userModule.NewProvider,

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
