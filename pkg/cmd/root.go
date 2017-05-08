package cmd

import (
	"github.com/spf13/cobra"

	appcli "bjdaos/pegasus/pkg/appoint/cli"
	apparatuscli "bjdaos/pegasus/pkg/instrument/heart/cli"
	pintocli "bjdaos/pegasus/pkg/pinto/cli"
	rptcli "bjdaos/pegasus/pkg/reporter/cli"
	wccli "bjdaos/pegasus/pkg/wc/cli"
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
		pegasus apparatus
	`,
}

func init() {
	RootCmd.AddCommand(wccli.NewWeiChatCmd("wc"))
	RootCmd.AddCommand(rptcli.NewReporterCmd("rpt"))
	RootCmd.AddCommand(appcli.NewAppointManagerCmd("app"))
	RootCmd.AddCommand(pintocli.NewPintoCmd("pinto"))
	RootCmd.AddCommand(apparatuscli.NewHeartDataCmd("apparatus"))
}
