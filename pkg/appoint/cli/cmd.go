package cli

import (
	"net/http"
	"os"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"bjdaos/pegasus/pkg/appoint/db"
	"bjdaos/pegasus/pkg/appoint/handler"
	"bjdaos/pegasus/pkg/appoint/model"

	"bjdaos/pegasus/pkg/appoint/appointment"
	org "bjdaos/pegasus/pkg/appoint/pinto"
)

func NewAppointManagerCmd(name string) *cobra.Command {
	wc := &cobra.Command{
		Use:   name,
		Short: "Manage appoint services",
	}

	wc.AddCommand(startCmd())
	return wc
}

func startCmd() *cobra.Command {
	var addr string
	var user, passwd, ip, port, dbname string
	var syncController org.Config
	start := &cobra.Command{
		Use:   "start",
		Short: "Start appointment system service",
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := db.Init(user, passwd, ip, port, dbname); err != nil {
				glog.Errorf("appointment init db err %v\n", err)
				os.Exit(1)
			}
			if err := model.Init(db.GetDB()); err != nil {
				glog.Errorf("appointment init model err %v\n", err)
				os.Exit(1)
			}

			appointment.Init()
			go syncController.Run()

		},
		Run: func(cmd *cobra.Command, args []string) {
			router := handler.CreateHttpRouter()
			if err := http.ListenAndServe(addr, router); err != nil {
				glog.Errorf("wc start err %v\n", err)
				os.Exit(1)
			}
		},
	}

	flags := start.Flags()
	flags.StringVar(&addr, "listen", ":9200", "TCP network address to listen on, to serve incomming http request.")
	flags.StringVar(&user, "db_user", "postgres", "Database user for the Application.")
	flags.StringVar(&passwd, "db_passwd", "postgres190@", "Database passwd for the Application.")
	flags.StringVar(&ip, "db_ip", "10.1.0.190", "Database ip for the Application.")
	flags.StringVar(&port, "db_port", "5432", "Database port for the Application.")
	flags.StringVar(&dbname, "db_name", "pinto", "Database name for the Application.")

	flags.StringVar(&syncController.User, "pinto_user", "postgres", "pinto user for the Application.")
	flags.StringVar(&syncController.Password, "pinto_passwd", "postgres190@", "pinto passwd for the Application.")
	flags.StringVar(&syncController.IP, "pinto_ip", "10.1.0.190", "pinto ip for the Application.")
	flags.StringVar(&syncController.Port, "pinto_port", "5432", "pinto port for the Application.")
	flags.StringVar(&syncController.Database, "pinto_name", "pinto", "pinto name for the Application.")

	return start
}
