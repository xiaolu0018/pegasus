package appoint

import (
	"192.168.199.199/bjdaos/pegasus/pkg/appoint/login"
	commonutil "192.168.199.199/bjdaos/pegasus/pkg/common/util"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/util"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const PASSWORD = "58f06cdfa46d12688c23405b"

func AuthUser(handler func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		user, pw, ok := r.BasicAuth()
		if !ok {
			unauthorizedHandler(w, r)
			return
		}

		loginuser, _ := login.Get(user)
		if loginuser != nil && loginuser.LoginAccount != "" {
			if commonutil.Md5([]byte(pw)) == loginuser.PassWord {
				ps = util.AddParam(ps, "user", user)
				handler(w, r, ps)
			} else {
				unauthorizedHandler(w, r)
			}
		} else { //这里是微信服务来的请求
			if pw == PASSWORD {
				ps = util.AddParam(ps, "user", user)
				handler(w, r, ps)

			} else {
				unauthorizedHandler(w, r)
			}
		}
	}
}

func unauthorizedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(401)
	w.Write([]byte(http.StatusText(401)))
	return
}
