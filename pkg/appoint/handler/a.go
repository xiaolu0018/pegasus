package handler

import (
	"bjdaos/pegasus/pkg/appoint/appointment"
	"bjdaos/pegasus/pkg/appoint/db"
	"bjdaos/pegasus/pkg/appoint/model"
	"bjdaos/pegasus/pkg/appoint/organization"
	org "bjdaos/pegasus/pkg/appoint/pinto"
	"flag"
	"github.com/golang/glog"
	"net/http"
	"os"
)

func Main_() {
	var addr string
	var user, passwd, ip, port, dbname string
	var syncController org.Config
	flag.StringVar(&addr, "listen", ":9200", "TCP network address to listen on, to serve incomming http request.")
	flag.StringVar(&user, "db_user", "postgres", "Database user for the Application.")
	flag.StringVar(&passwd, "db_passwd", "postgres190@", "Database passwd for the Application.")
	flag.StringVar(&ip, "db_ip", "10.1.0.190", "Database ip for the Application.")
	flag.StringVar(&port, "db_port", "5432", "Database port for the Application.")
	flag.StringVar(&dbname, "db_name", "pinto", "Database name for the Application.")

	flag.StringVar(&syncController.User, "pinto_user", "postgres", "pinto user for the Application.")
	flag.StringVar(&syncController.Password, "pinto_passwd", "postgres190@", "pinto passwd for the Application.")
	flag.StringVar(&syncController.IP, "pinto_ip", "10.1.0.190", "pinto ip for the Application.")
	flag.StringVar(&syncController.Port, "pinto_port", "5432", "pinto port for the Application.")
	flag.StringVar(&syncController.Database, "pinto_name", "pinto", "pinto name for the Application.")
	if err := db.Init(user, passwd, ip, port, dbname); err != nil {
		glog.Errorf("appointment init db err %v\n", err)
		os.Exit(1)
	}
	if err := model.Init(db.GetDB()); err != nil {
		glog.Errorf("appointment init model err %v\n", err)
		os.Exit(1)
	}

	go appointment.StartController()
	organization.Init()
	go syncController.Run()
	router := CreateHttpRouter()
	if err := http.ListenAndServe(addr, router); err != nil {
		glog.Errorf("wc start err %v\n", err)
		os.Exit(1)
	}

}
