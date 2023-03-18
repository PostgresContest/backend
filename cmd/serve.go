package cmd

import (
	"github.com/spf13/cobra"
)

func serveCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start server",
		Run: func(cmd *cobra.Command, args []string) {
			// todo
		},
	}
}
