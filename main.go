package main

import (
	"bjdaos/pegasus/pkg/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
