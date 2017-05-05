package login

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"

	httputil "github.com/1851616111/util/http"

	"bjdaos/pegasus/pkg/common/util/md5"
)

func CheckLoginHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()
	login := r.FormValue("login")
	password := r.FormValue("password")
	ok, err := CheckLogin(login, password)
	if err != nil || !ok {
		httputil.Response(rw, 401, err)
		return
	}
	httputil.Response(rw, 200, nil)
}

func ChangeLoginHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()
	login := r.FormValue("login")
	newpassword := r.FormValue("newpassword")
	oldpasswod := r.FormValue("oldpassword")
	ok, err := CheckLogin(login, oldpasswod)
	if err != nil || !ok {
		httputil.Response(rw, 401, err)
		return
	}
	if err := ChangePWD(login, newpassword); err != nil {
		httputil.Response(rw, 400, "change password err")
	}
	httputil.Response(rw, 200, nil)
}

func CheckLogin(loginaccount, password string) (bool, error) {
	loginuser, err := Get(loginaccount)
	if err != nil {
		glog.Errorln("login.CheckLogin err ", err)
		return false, err
	}
	if md5.Md5([]byte(password)) == loginuser.PassWord {
		return true, nil
	}
	return false, nil
}
