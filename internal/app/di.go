package app

import (
	"backend/internal/infrastructure/auth"
	"backend/internal/infrastructure/config"
	dbPrivate "backend/internal/infrastructure/db/private"
	dbPublic "backend/internal/infrastructure/db/public"
	"backend/internal/infrastructure/executor"
	openapiV1 "backend/internal/infrastructure/handlers/openapi/v1"
	authModule "backend/internal/infrastructure/handlers/openapi/v1/modules/auth"
	taskModule "backend/internal/infrastructure/handlers/openapi/v1/modules/task"
	userModule "backend/internal/infrastructure/handlers/openapi/v1/modules/user"
	attemptRepository "backend/internal/infrastructure/repositories/attempt"
	queryRepository "backend/internal/infrastructure/repositories/query"
	taskRepository "backend/internal/infrastructure/repositories/task"
	userRepository "backend/internal/infrastructure/repositories/user"
	"backend/internal/infrastructure/server"
	"backend/internal/logger"

	"go.uber.org/fx"
)

func getProvidersAndInvokers() ([]any, []any) {
	providers := []any{
		config.NewProvider,
		logger.NewProvider,

		dbPrivate.NewProvider,
		dbPublic.NewProvider,
		auth.NewAccessTokenProvider,
		auth.NewSecurityProvider,

		userRepository.NewProvider,
		queryRepository.NewProvider,
		taskRepository.NewProvider,
		attemptRepository.NewProvider,
		executor.NewProvider,

		authModule.NewProvider,
		userModule.NewProvider,
		taskModule.NewProvider,

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
