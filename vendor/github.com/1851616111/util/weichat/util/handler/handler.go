package handler

import (
	"github.com/1851616111/util/weichat/util/sign"
	token "github.com/1851616111/util/weichat/util/user-token"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var APP_ID string
var Token *token.Config

//validate qualification of weichat developer
func DeveloperValidater(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	s := sign.Sign(r.FormValue("nonce"), r.FormValue("timestamp"), APP_ID)

	if s != r.FormValue("signature") {
		w.WriteHeader(400)
	} else {
		w.WriteHeader(200)
		w.Write([]byte(r.FormValue("echostr")))
	}

	return
}

func AuthValidator(tokenCallBack func(*token.Token, *httprouter.Params) error, handler func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		r.ParseForm()
		code := r.FormValue("code")
		glog.Errorf("authValidater get user code=%s\n", code)

		if code == "" {
			forbiddenHandler(w, r, ps)
			return
		}

		tk, err := Token.Exchange(code)
		if err != nil {
			glog.Errorf("authValidater exchange access_token err %v\n", err)
			forbiddenHandler(w, r, ps)
			return
		}
		glog.Errorf("authValidater exchange access_token %v\n", *tk)

		if err := tokenCallBack(tk, &ps); err != nil {
			glog.Errorf("authValidater tokenCallBack(%v) err %v\n", tk.Open_ID, err)
			forbiddenHandler(w, r, ps)
			return
		}

		handler(w, r, ps)
		return
	}
}

func forbiddenHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte(`<html><head><title>迪安</title></head><body>请在微信客户端打开链接</body></html>`))
	return
}
