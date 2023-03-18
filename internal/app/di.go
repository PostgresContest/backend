package app

import (
	"backend/internal/config"
	"backend/internal/handlers/openapi/v1/modules/auth"
	"backend/internal/logger"
	"backend/internal/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func getProvidersAndInvokers() ([]any, []any) {
	providers := []any{
		config.NewProvider,
		logger.NewProvider,

		auth.NewProvider,

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
		fx.WithLogger(func(log *logger.Logger) fxevent.Logger {
			return log
		}),
	)
	return f
}
