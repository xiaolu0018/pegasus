package cmd

import (
	"github.com/spf13/cobra"

	appcli "192.168.199.199/bjdaos/pegasus/pkg/appoint/cli"
	pintocli "192.168.199.199/bjdaos/pegasus/pkg/pinto/cli"
	rptcli "192.168.199.199/bjdaos/pegasus/pkg/reporter/cli"
	wccli "192.168.199.199/bjdaos/pegasus/pkg/wc/cli"
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
		pegasus pinto
	`,
}

func init() {
	RootCmd.AddCommand(wccli.NewWeiChatCmd("wc"))
	RootCmd.AddCommand(rptcli.NewReporterCmd("rpt"))
	RootCmd.AddCommand(appcli.NewAppointManagerCmd("app"))
	RootCmd.AddCommand(pintocli.NewPintoCmd("pinto"))
}
