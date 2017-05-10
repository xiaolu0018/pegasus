package cli

import (
	"os"

	"net/http"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"bjdaos/pegasus/pkg/pinto/server/db"
	"bjdaos/pegasus/pkg/pinto/server/handler"
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
	var w_user, w_passwd, w_ip, w_port, w_dbname string
	var r_user, r_passwd, r_ip, r_port, r_dbname string
	start := &cobra.Command{
		Use:   "start",
		Short: "Start appointment system service",
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := db.InitReadDB(r_user, r_passwd, r_ip, r_port, r_dbname); err != nil {
				glog.Errorf("init read database err %v\n", err)
				os.Exit(1)
			}
			if err := db.InitWriteDB(w_user, w_passwd, w_ip, w_port, w_dbname); err != nil {
				glog.Errorf("init read database err %v\n", err)
				os.Exit(1)
			}
		},

		Run: func(cmd *cobra.Command, args []string) {
			glog.Warningln("pinto Start")
			router := handler.CreateHttpRouter()
			if err := http.ListenAndServe(addr, router); err != nil {
				glog.Errorf("pinto start err %v\n", err)
				os.Exit(1)
			}
		},
	}

	flags := start.Flags()
	flags.StringVar(&addr, "listen", ":9300", "TCP network address to listen on, to serve incomming http request.")
	flags.StringVar(&r_user, "r_db_user", "postgres", "Database user for the Application.")
	flags.StringVar(&r_passwd, "r_db_passwd", "postgres190@", "Database passwd for the Application.")
	flags.StringVar(&r_ip, "r_db_ip", "10.1.0.190", "Database ip for the Application.")
	flags.StringVar(&r_port, "r_db_port", "5432", "Database port for the Application.")
	flags.StringVar(&r_dbname, "r_db_name", "pinto", "Database name for the Application.")

	flags.StringVar(&w_user, "w_db_user", "postgres", "Database user for the Application.")
	flags.StringVar(&w_passwd, "w_db_passwd", "postgres190@", "Database passwd for the Application.")
	flags.StringVar(&w_ip, "w_db_ip", "10.1.0.190", "Database ip for the Application.")
	flags.StringVar(&w_port, "w_db_port", "5432", "Database port for the Application.")
	flags.StringVar(&w_dbname, "w_db_name", "pinto", "Database name for the Application.")

	return start
}
