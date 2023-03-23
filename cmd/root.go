package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	root := &cobra.Command{
		Use: "postgres-contest",
		Run: func(cmd *cobra.Command, args []string) {
			_, _ = os.Stdout.WriteString("No commands provided\n")
		},
	}

	root.AddCommand(serveCommand())
	root.AddCommand(dbDsnCommand())

	return root
}
