package main

import (
	"os"

	"backend/cmd"

	_ "github.com/lib/pq"
)

func main() {
	if err := cmd.Root().Execute(); err != nil {
		_, _ = os.Stdout.WriteString(err.Error())
		os.Exit(1)
	}
}
