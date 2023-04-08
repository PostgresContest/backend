package cmd

import (
	"os"

	"backend/internal/infrastructure/config"
	"backend/internal/infrastructure/db/private"
	"backend/internal/logger"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func dbDsnCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dsn",
		Short: "Get database dsn",
		Run: func(cmd *cobra.Command, args []string) {
			_, _ = os.Stdout.WriteString("Which database dsn I need to retrieve? Run with \"--help\"\n")
		},
	}

	providers := fx.Provide(
		config.NewProvider,
		logger.NewProviderWithDiscardOutput,
		private.NewProvider,
	)

	privateDsn := &cobra.Command{
		Use:   "private",
		Short: "Retrieves private database dsn",
		Run: func(cmd *cobra.Command, args []string) {
			dsn := ""
			a := fx.New(
				providers,
				fx.Invoke(func(connection *private.Connection, sd fx.Shutdowner) {
					dsn = connection.Dsn
					_ = sd.Shutdown()
				}),
				fx.NopLogger,
			)
			a.Run()
			_, _ = os.Stdout.WriteString(dsn + "\n")
		},
	}

	cmd.AddCommand(privateDsn)

	return cmd
}
