package cli

import (
	"os"

	"net/http"
	"path/filepath"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"bjdaos/pegasus/pkg/wc/db"
	"bjdaos/pegasus/pkg/wc/user"

	"bjdaos/pegasus/pkg/wc/handler"
	"bjdaos/pegasus/pkg/wc/manager"

	"bjdaos/pegasus/pkg/wc/server/start"
	"github.com/julienschmidt/httprouter"
)

func NewWeiChatCmd(name string) *cobra.Command {
	wc := &cobra.Command{
		Use:   name,
		Short: "Manage weichat services",
	}

	wc.AddCommand(startCmd())
	wc.AddCommand(startActivityCmd())
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

			if err := handler.InitApiToken(appID, appSecret); err != nil {
				glog.Errorf("wc start: init wc api token controller err %v\n", err)
				os.Exit(1)
			}

			if err := manager.CreateMenu(); err != nil {
				glog.Errorf("wc start:init menu err %v", err)
				os.Exit(1)
			}
		},

		Run: func(cmd *cobra.Command, args []string) {
			if err := user.Init(); err != nil {
				glog.Errorf("wc init user cache err %v\n", err)
				os.Exit(1)
			}

			manager, err := handler.NewMenuRedirectManager(base, "bear_token")
			if err != nil {
				glog.Errorf("wc create redirect manager err %v\n", err)
				os.Exit(1)
			}
			manager.Redirect("/api/appoint", "regist.html", handler.CompleteAccessTokenInfo)
			manager.Redirect("/api/branch", "branch.html", handler.CompleteAccessTokenInfo)
			manager.Redirect("/api/reportmenu", "myRep.html", handler.CompleteAccessTokenInfo)

			router := httprouter.New()

			handler.AddApiToRouter(router, dist)
			if err := manager.AddRedirectToRouter(router); err != nil {
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

	flags.StringVar(&addr, "listen", ":9001", "TCP network address to listen on, to serve incomming http request.")
	//flags.StringVar(&appSecret, "name", "bjdaos", "The Weichat Official Accounts app name, must meed to set")
	flags.StringVar(&appID, "id", "wxd09c7682905819e6", "The Weichat Official Accounts app id, must need to set")
	flags.StringVar(&appSecret, "secret", "b9938ddfec045280eba89fab597a0c41", "The Weichat Official Accounts app secret, must meed to set")
	flags.StringVar(&base, "base", "http://www.elepick.com/dist", "The app base home url")
	flags.StringVar(&dist, "dist", "./dist", "The file path one application server.")

	return start
}

func startActivityCmd() *cobra.Command {
	o := start.NewActivityConfig()

	var start = &cobra.Command{
		Use:   "start-activity",
		Short: "Start weichat activity internet service",
		PreRun: func(cmd *cobra.Command, args []string) {
		},

		Run: func(cmd *cobra.Command, args []string) {
			var router *httprouter.Router = httprouter.New()
			if err := o.Start(router); err != nil {
				glog.Errorf("wc start activity err %v\n", err)
				os.Exit(1)
			}

			if err := http.ListenAndServe(o.ListenAddress, router); err != nil {
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
	flags.StringVar(&o.Schema, "schema", "http", "The schema for http server")
	flags.StringVar(&o.Domain, "domain", "hd1.dahe100.cn", "The activity server domain")
	flags.StringVar(&o.DistBasePath, "basepath", "dist", "The activity server static file basepath")
	flags.StringVar(&o.LocalDistPath, "dist", "./dist/activity", "The file path one application server.")

	flags.StringVar(&o.ListenAddress, "listen", ":80", "TCP network address to listen on, to serve incomming http request.")

	flags.StringVar(&o.AppSignToken, "token", "123456", "The Weichat Official Accounts app developer api token, must need to set")
	//flags.StringVar(&o.AppID, "id", "wxd38ad1e4f070b5e6", "The Weichat Official Accounts app id, must need to set")
	//flags.StringVar(&o.AppSecret, "secret", "628c1bdd5f8b44ebfd9ef2e28afe38ea", "The Weichat Official Accounts app secret, must meed to set")
	//
	flags.StringVar(&o.AppID, "id", "wxadb8865cdb995ded", "The Weichat Official Accounts app id, must need to set")
	flags.StringVar(&o.AppSecret, "secret", "2d5ac69b6b76dbde88e71b3663e41ff4", "The Weichat Official Accounts app secret, must meed to set")

	flags.StringVar(&o.User, "user", "wxhd", "Database user for the Application.")
	flags.StringVar(&o.Password, "passwd", "wXHd2069", "Database passwd for the Application.")
	flags.StringVar(&o.Ip, "ip", "10.0.108.21", "Database ip for the Application.")
	flags.StringVar(&o.Port, "port", "5432", "Database port for the Application.")
	flags.StringVar(&o.Database, "name", "wxhd", "Database name for the Application.")

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
