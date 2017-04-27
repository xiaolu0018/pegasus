package cli

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/cobra"

	"192.168.199.199/bjdaos/pegasus/pkg/wc/db"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/handler"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/manager"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/token"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/user"
)

func NewWeiChatCmd(name string) *cobra.Command {
	wc := &cobra.Command{
		Use:   name,
		Short: "Manage weichat services",
	}

	wc.AddCommand(startCmd())
	return wc
}

func startCmd() *cobra.Command {
	var addr string
	var appID, appSecret string
	var base string
	var dist string
	var dbuser, passwd, ip, port, dbname string
	var start = &cobra.Command{
		Use:   "start",
		Short: "Start weichat internet service",
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := validateDist(dist); err != nil {
				glog.Errorf("wc start: validate dist path err %v\n", err)
				os.Exit(1)
			}

			if err := db.Init(dbuser, passwd, ip, port, dbname); err != nil {
				glog.Errorf("appointment init db err %v\n", err)
				os.Exit(1)
			}

			if err := token.InitApiToken(appID, appSecret); err != nil {
				glog.Errorf("wc start: init wc api token controller err %v\n", err)
				os.Exit(1)
			}

			if err := manager.CreateMenu(token.TokenCtrl.GetToken()); err != nil {
				glog.Errorf("wc start:init menu err", err)
				os.Exit(1)
			}

		},

		Run: func(cmd *cobra.Command, args []string) {
			if err := user.Init(); err != nil {
				glog.Errorf("wc init user cache err %v\n", err)
				os.Exit(1)
			}

			manager, err := handler.NewRedirectManager(base)
			if err != nil {
				glog.Errorf("wc create redirect manager err %v\n", err)
				os.Exit(1)
			}
			manager.Redirect("/api/appoint", "regist.html")
			manager.Redirect("/api/branch", "branch.html")
			manager.Redirect("/api/reportmenu", "myRep.html")

			router := handler.CreateHttpRouter(dist)
			if err := manager.AddRouter(router.(*httprouter.Router)); err != nil {
				glog.Errorf("wc add redirect to router err %v\n", err)
				os.Exit(1)
			}

			if err := http.ListenAndServe(addr, router); err != nil {
				glog.Errorf("wc start err %v\n", err)
				os.Exit(1)
			}
		},
	}

	//APP_NAME = "bjdaos"
	//APP_ID    = "wxd09c7682905819e6"
	//APP_SECRET   = "b9938ddfec045280eba89fab597a0c41"
	//IndexUrl = "http://www.elepick.com:8080/pegasus/dist"

	flags := start.Flags()

	flags.StringVar(&dbuser, "db_user", "postgres", "Database user for the Application.")
	flags.StringVar(&passwd, "db_passwd", "postgres190@", "Database passwd for the Application.")
	flags.StringVar(&ip, "db_ip", "10.1.0.190", "Database ip for the Application.")
	flags.StringVar(&port, "db_port", "5432", "Database port for the Application.")
	flags.StringVar(&dbname, "db_name", "pinto", "Database name for the Application.")

	flags.StringVar(&addr, "listen", ":9000", "TCP network address to listen on, to serve incomming http request.")
	flags.StringVar(&appSecret, "name", "bjdaos", "The Weichat Official Accounts app name, must meed to set")
	flags.StringVar(&appID, "id", "wxd09c7682905819e6", "The Weichat Official Accounts app id, must need to set")
	flags.StringVar(&appSecret, "secret", "b9938ddfec045280eba89fab597a0c41", "The Weichat Official Accounts app secret, must meed to set")
	flags.StringVar(&base, "base", "http://www.elepick.com/dist", "The app base home url")
	flags.StringVar(&dist, "dist", "./dist", "The file path one application server.")
	return start
}

func validateDist(dist string) (err error) {
	dist, err = filepath.Abs(dist)
	if err != nil {
		return
	}

	_, err = os.Stat(dist)
	return
}
