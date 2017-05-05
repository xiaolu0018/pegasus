package cli

import (
	"net/http"
	"os"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	//"bjdaos/pegasus/pkg/common/util/safe"
	"bjdaos/pegasus/pkg/reporter/db"
	"bjdaos/pegasus/pkg/reporter/handler"
	"bjdaos/pegasus/pkg/reporter/model"
)

func NewReporterCmd(name string) *cobra.Command {
	wc := &cobra.Command{
		Use:   name,
		Short: "Manage reporter services",
	}

	wc.AddCommand(startCmd())
	return wc
}

func startCmd() *cobra.Command {
	var addr string
	var w_user, w_passwd, w_ip, w_port, w_dbname string
	var user, passwd, ip, port, dbname string
	//var publicKeyPath string
	start := &cobra.Command{
		Use:   "start",
		Short: "Start reporter system service",
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := db.InitFunction(w_user, w_passwd, w_ip, w_port, w_dbname); err != nil {
				glog.Errorf("reporter init db err %v\n", err)
				os.Exit(1)
			}

			if err := db.Init(user, passwd, ip, port, dbname); err != nil {
				glog.Errorf("reporter init db err %v\n", err)
				os.Exit(1)
			}
			if err := model.Init(db.GetReadDB()); err != nil {
				glog.Errorf("reporter init model err %v\n", err)
				os.Exit(1)
			}

			//if err := safe.LoadPubKey(publicKeyPath); err != nil {
			//	glog.Errorf("reporter init rsa err %v\n", err)
			//	os.Exit(1)
			//}
		},

		Run: func(cmd *cobra.Command, args []string) {
			router := handler.CreateHttpRouter()
			if err := http.ListenAndServe(addr, router); err != nil {
				glog.Errorf("reporter start err %v\n", err)
				os.Exit(1)
			}
		},
	}

	//export DB_USER=postgres
	//export DB_PASSWD=postgresql2016
	//export DB_IP=192.168.199.216
	//export DB_PORT=5432
	//export DB_NAME=pinto
	//#export DB_NAME=mat
	//dbEnv = env.NewEnv("PostgresSQL", "DB_USER", "DB_PASSWD", "DB_IP", "DB_PORT", "DB_NAME")

	flags := start.Flags()
	flags.StringVar(&addr, "listen", ":9100", "TCP network address to listen on, to serve incomming http request.")
	flags.StringVar(&user, "r_db_user", "postgres", "Database user for the Application.")
	flags.StringVar(&passwd, "r_db_passwd", "postgres190@", "Database passwd for the Application.")
	flags.StringVar(&ip, "r_db_ip", "10.1.0.190", "Database ip for the Application.")
	flags.StringVar(&port, "r_db_port", "5432", "Database port for the Application.")
	flags.StringVar(&dbname, "r_db_name", "pinto", "Database name for the Application.")
	//flags.StringVar(&publicKeyPath, "public_key", "public.pem", "Transport layer data encoding")

	flags.StringVar(&w_user, "w_db_user", "postgres", "Database user for the Application.")
	flags.StringVar(&w_passwd, "w_db_passwd", "postgres190@", "Database passwd for the Application.")
	flags.StringVar(&w_ip, "w_db_ip", "10.1.0.190", "Database ip for the Application.")
	flags.StringVar(&w_port, "w_db_port", "5432", "Database port for the Application.")
	flags.StringVar(&w_dbname, "w_db_name", "pinto", "Database name for the Application.")

	return start
}
