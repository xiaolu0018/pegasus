package cmd

import (
	"github.com/spf13/cobra"
	"bjdaos/pegasus/pkg/apparatus/electrocardiogram/cli"
)

var RootCmd = &cobra.Command{
	Use: "bjdaos_tool",
	Short: "bjdaos_tool App cmd",
	Long: `bjdaos_tool is a tool to manage golang servic`,
	Example:`
		bjdaos_tool hd
	`,
}

func init(){
	RootCmd.AddCommand(cli.HeartDataCmd("hd"))
}