package cmd

import (
	"backend/internal/app"

	"github.com/spf13/cobra"
)

func serveCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start server",
		Run: func(cmd *cobra.Command, args []string) {
			a := app.NewApp()
			a.Start()
		},
	}
}
