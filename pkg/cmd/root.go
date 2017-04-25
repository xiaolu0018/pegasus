package cmd

import (
	appcli "192.168.199.199/bjdaos/pegasus/pkg/appoint/cli"
	rptcli "192.168.199.199/bjdaos/pegasus/pkg/reporter/cli"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/cli"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "pegasus",
	Short: "pegasus application commands",
	Long:  `pegasus is a tool to mamage golang service.`,
	Example: `
		pegasus wc
		pegasus rpt
		pegasus app
	`,
}

func init() {
	RootCmd.AddCommand(cli.NewWeiChatCmd("wc"))
	RootCmd.AddCommand(rptcli.NewReporterCmd("rpt"))
	RootCmd.AddCommand(appcli.NewAppointManagerCmd("app"))
}
