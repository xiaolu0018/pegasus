package handler

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"

	"192.168.199.199/bjdaos/pegasus/pkg/wc/user"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/util"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/common"
)

func authUser(handler func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		glog.Errorln("authuser", isTokenInValid(r))
		if isTokenInValid(r) {
			unauthorizedHandler(w, r)
			return
		}

		token := r.Header[user.TokenHeaderName][0]
		ok, id := user.IDCache.Auth(token)
		if !ok {
			unauthorizedHandler(w, r)
			return
		}

		ps = util.AddParam(ps, common.AuthHeaderKey, id)
		handler(w, r, ps)
	}
}

func authAdmin(handler func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		user, passwd, ok := r.BasicAuth()
		if !ok {
			unauthorizedHandler(w, r)
			return
		}
		if user != "admin" && passwd != "123456" {
			forbiddenHandler(w, r)
			return
		}
		handler(w, r, ps)
		return
	}
}

func unauthorizedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(401)
	w.Write([]byte(http.StatusText(401)))
	return
}

func forbiddenHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(403)
	w.Write([]byte(http.StatusText(403)))
	return
}

func isTokenInValid(r *http.Request) bool {
	token := r.Header.Get(user.TokenHeaderName)
	return len(token) != user.TokenLength
}
