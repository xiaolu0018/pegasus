package model

import (
	"errors"
	"fmt"
	"github.com/golang/glog"
)

func Auth(username string, password string) (string, error) {
	if len(username) > 20 || len(password) > 50 {
		return "", errors.New("invalid username or password")
	}

	sql := fmt.Sprintf(`select (SELECT count(*) FROM manager where account='%s' and password='%s') as num, hos_code FROM manager where account='%s' and password ='%s'`,
		username, password, username, password)

	var count int
	var hosCode string
	row := DB.QueryRow(sql)
	err := row.Scan(&count, &hosCode)
	if err != nil {
		glog.Errorf("GetLoginInfo: sql return err %v\n", err)
		return "", err
	}
	glog.Errorln("hostcode count", hosCode, count)
	if count == 1 {
		return hosCode, nil
	} else {
		return "", errors.New("user not found")
	}
}

func GetHosCode(username string) (*string, error) {

	var hos_code *string
	row := DB.QueryRow(fmt.Sprintf("SELECT hos_code FROM manager WHERE account='%s'", username))
	err := row.Scan(&hos_code)
	if err != nil {
		glog.Errorf("GetHosCode: sql return err %v\n", err)
		return nil, err
	}

	return hos_code, err
}
