package app

import (
	"backend/internal/config"
	"backend/internal/logger"
	"go.uber.org/fx"
)

func getProvidersAndInvokers() ([]any, []any) {
	providers := []any{
		config.NewProvider,
		logger.NewProvider,
	}
	invokers := []any{
		func(log logger.Logger) {
			log.Info("Wow")
		},
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
