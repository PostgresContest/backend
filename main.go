package main

import (
	"backend/cmd"
	"log"
)

func main() {
	if err := cmd.Root().Execute(); err != nil {
		log.Fatal(err)
	}
}
