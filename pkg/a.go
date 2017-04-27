package appoint

import (
	"192.168.199.199/bjdaos/pegasus/pkg/appoint/db"
	"192.168.199.199/bjdaos/pegasus/pkg/appoint/handler"
	"192.168.199.199/bjdaos/pegasus/pkg/appoint/model"
	"errors"
	"github.com/golang/glog"
	"net/http"
	"os"
	"strings"
)

func Main_() {
	var user, passwd, ip, port, dbname string
	user = "postgres"
	passwd = "postgres190@"
	ip = "10.1.0.190"
	port = "5432"
	dbname = "pinto"
	base := "http://127.0.0.1:3000/backstages"

	if err := db.Init(user, passwd, ip, port, dbname); err != nil {
		glog.Errorf("appointment init db err %v\n", err)
		os.Exit(1)
	}
	if err := model.Init(db.GetDB()); err != nil {
		glog.Errorf("appointment init model err %v\n", err)
		os.Exit(1)
	}
	manager, err := NewRedirectManager(base)
	if err != nil {
		glog.Errorf("wc create redirect manager err %v\n", err)
		os.Exit(1)
	}
	manager.Redirect("/api/index", "index.html")
	router := handler.CreateHttpRouter()
	if err := http.ListenAndServe(":3000", router); err != nil {
		glog.Errorf("wc start err %v\n", err)
		os.Exit(1)
	}

}

type redirectManager struct {
	baseUrl                string
	pathToResourceMappings map[string]string //重定向路径与重定向资源映射
}

func NewRedirectManager(base string) (*redirectManager, error) {
	if base == "" {
		return nil, errors.New("redirect basepath not found")
	}

	if strings.LastIndex(base, string(os.PathSeparator)) == len(base)-1 {
		base = base[:len(base)-1]
	}

	return &redirectManager{
		baseUrl:                base,
		pathToResourceMappings: map[string]string{},
	}, nil
}

func (m *redirectManager) Redirect(path, Resource string) {
	m.pathToResourceMappings[path] = Resource
}
