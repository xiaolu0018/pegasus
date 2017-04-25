package main

import (
	"192.168.199.199/bjdaos/pegasus/pkg/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
