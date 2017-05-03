package cli

import "fmt"

type activityOption struct {
	schema       string
	domain       string
	distBasePath string

	dist          string
	appID         string
	appSecret     string
	listenAddress string

	db_user     string
	db_password string
	db_ip       string
	db_port     string
	db_name     string
}

func (o *activityOption) GetAbsPath() string {
	return fmt.Sprintf("%s://%s/%s", o.schema, o.domain, o.distBasePath)
}
