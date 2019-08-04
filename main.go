package main

import (
	"os"

	"github.com/radoondas/earthquakebeat/cmd"

	_ "github.com/radoondas/earthquakebeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
