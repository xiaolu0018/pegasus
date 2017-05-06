//==================================================================
//创建时间：2017-4-27 首次创建
//功能描述：Flag实现配置
//创建人：郭世江
//修改记录：若要修改请记录
//==================================================================
package cli

import (
	"os"

	"net/http"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"bjdaos/pegasus/pkg/apparatus/electrocardiogram/api"
	"bjdaos/pegasus/pkg/apparatus/electrocardiogram/db"
	"bjdaos/pegasus/pkg/apparatus/electrocardiogram/handler"
	"bjdaos/pegasus/pkg/apparatus/electrocardiogram/model"
)

func NewHeartDataCmd(name string) *cobra.Command {
	bjdaos := &cobra.Command{
		Use:   name,
		Short: "heartdata services",
	}
	bjdaos.AddCommand(startCmd())
	return bjdaos
}

func startCmd() *cobra.Command {
	var addr string
	var user, passwd, ip, port, dbname string
	start := &cobra.Command{
		Use:   "start",
		Short: "Start heart_tdata system service",
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := db.Init(user, passwd, ip, port, dbname); err != nil {
				glog.Errorf("reporter init db err %v\n", err)
				os.Exit(1)
			}
			if err := model.Init(db.GetDB()); err != nil {
				glog.Errorf("reporter init model err %v\n", err)
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			router := handler.CreateHttpRouter()
			if err := http.ListenAndServe(addr, router); err != nil {
				glog.Errorf("reporter start err %v\n", err)
				os.Exit(1)
			}
		},
	}
	flags := start.Flags()
	flags.StringVar(&addr, "listen", ":9050", " http request listen port")
	flags.StringVar(&user, "db_user", "postgres", "App Database User")
	flags.StringVar(&passwd, "db_passwd", "postgres190@", "App Database Passwd")
	flags.StringVar(&ip, "db_ip", "10.1.0.190", "App Database IP")
	flags.StringVar(&port, "db_port", "5432", "App Database Port")
	flags.StringVar(&dbname, "db_name", "pinto", "App Database Name")
	flags.StringVar(&api.BaseUrl, "baseurl", "http://10.1.0.190:8080", "serve address")
	flags.StringVar(&handler.ImgsPathUrl, "img_path", "/IdeaProjects/src/xml/", "ftp server storage path")
	flags.StringVar(&handler.XmlPathUrl, "xmp_path", "../xml/", "xml file storage path")
	return start
}
