package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	root := &cobra.Command{
		Use: "postgres-contest",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("No commands provided")
		},
	}

	root.AddCommand(serveCommand())

	return root
}
