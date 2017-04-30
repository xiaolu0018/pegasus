package login

import (
	"192.168.199.199/bjdaos/pegasus/pkg/common/util"
	"192.168.199.199/bjdaos/pegasus/pkg/reporter/db"
	"fmt"
	"github.com/golang/glog"
)

func Get(account string) (*LoginUser, error) {
	sqlStr := fmt.Sprintf("SELECT loginname,org_code,password FROM %s WHERE loginaccount = '%s'", TABLE_LOGINUSER, account)
	loginUser := LoginUser{}
	if err := db.GetReadDB().QueryRow(sqlStr).Scan(&loginUser.LoginName, &loginUser.OrgCode, &loginUser.PassWord); err != nil {
		glog.Errorln("login.Get err ", err)
		return nil, err
	}
	loginUser.LoginAccount = account
	return &loginUser, nil
}

func ChangePWD(account, newpwd string) error {

	newpwd = util.Md5([]byte(newpwd))
	sqlStr := fmt.Sprintf("UPDATE TABLE %s SET password = '%s' WHERR loginaccount = '%d'", TABLE_LOGINUSER, newpwd, account)
	if _, err := db.GetReadDB().Exec(sqlStr); err != nil {
		glog.Errorln("login.ChangePED err ", err)
		return err
	}
	return nil
}
