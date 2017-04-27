package cli

import (

	"github.com/spf13/cobra"
)

func NewPintoCmd(name string) *cobra.Command {
	wc := &cobra.Command{
		Use:   name,
		Short: "Manage pinto services",
	}

	wc.AddCommand(startCmd())
	return wc
}

func startCmd() *cobra.Command {
	var addr string
	var user, passwd, ip, port, dbname string
	start := &cobra.Command{
		Use:   "start",
		Short: "Start appointment system service",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	flags := start.Flags()
	flags.StringVar(&addr, "listen", ":9300", "TCP network address to listen on, to serve incomming http request.")
	flags.StringVar(&user, "r_db_user", "postgres", "Database user for the Application.")
	flags.StringVar(&passwd, "r_db_passwd", "postgres190@", "Database passwd for the Application.")
	flags.StringVar(&ip, "r_db_ip", "10.1.0.190", "Database ip for the Application.")
	flags.StringVar(&port, "r_db_port", "5432", "Database port for the Application.")
	flags.StringVar(&dbname, "r_db_name", "pinto", "Database name for the Application.")

	flags.StringVar(&user, "w_db_user", "postgres", "Database user for the Application.")
	flags.StringVar(&passwd, "w_db_passwd", "postgres190@", "Database passwd for the Application.")
	flags.StringVar(&ip, "w_db_ip", "10.1.0.190", "Database ip for the Application.")
	flags.StringVar(&port, "w_db_port", "5432", "Database port for the Application.")
	flags.StringVar(&dbname, "w_db_name", "pinto", "Database name for the Application.")

	return start
}
