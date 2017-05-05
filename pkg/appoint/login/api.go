package login

import (
	"fmt"
	"github.com/golang/glog"

	"bjdaos/pegasus/pkg/appoint/db"
	"bjdaos/pegasus/pkg/common/util/md5"
)

func Get(account string) (*LoginUser, error) {
	sqlStr := fmt.Sprintf("SELECT loginname,org_code,password FROM %s WHERE loginaccount = '%s'", TABLE_LOGINUSER, account)
	loginUser := LoginUser{}
	if err := db.GetDB().QueryRow(sqlStr).Scan(&loginUser.LoginName, &loginUser.OrgCode, &loginUser.PassWord); err != nil {
		glog.Errorln("login.Get err ", err)
		return nil, err
	}
	loginUser.LoginAccount = account
	return &loginUser, nil
}

func ChangePWD(account, newpwd string) error {

	newpwd = md5.Md5([]byte(newpwd))
	sqlStr := fmt.Sprintf("UPDATE TABLE %s SET password = '%s' WHERR loginaccount = '%d'", TABLE_LOGINUSER, newpwd, account)
	if _, err := db.GetDB().Exec(sqlStr); err != nil {
		glog.Errorln("login.ChangePED err ", err)
		return err
	}
	return nil
}
