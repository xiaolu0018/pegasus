package appoint

import (
	"192.168.199.199/bjdaos/pegasus/pkg/wc/util"

	"github.com/golang/glog"
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
		if pw == PASSWORD {
			glog.Errorln("user", user)
			ps = util.AddParam(ps, "user", user)
			handler(w, r, ps)
		} else {
			unauthorizedHandler(w, r)
		}
	}
}

func unauthorizedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(401)
	w.Write([]byte(http.StatusText(401)))
	return
}
