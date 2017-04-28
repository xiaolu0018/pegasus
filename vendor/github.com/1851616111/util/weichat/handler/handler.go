package handler

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"sync"

	httputil "github.com/1851616111/util/http"
	"github.com/1851616111/util/weichat/event"
	"github.com/1851616111/util/weichat/util/sign"
	token "github.com/1851616111/util/weichat/util/user-token"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
)

var APP_ID string
var Token *token.Config
var EventManager *event.EventManager
var EOnceL sync.Once

func init() {
	EOnceL.Do(func() {
		EventManager = event.NewEventManager()
	})
}

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

func EventAction(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	xmlData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputil.Response(w, 400, err)
		glog.Errorf("read weichat notify event err %v\n", err)
		return
	}

	e := event.Event{}
	if err := xml.Unmarshal(xmlData, &e); err != nil {
		httputil.Response(w, 400, err)
		glog.Errorf("decode weichat notify event msg err %v\n", err)
		return
	}

	if act := EventManager.Handle(&e); act != nil {
		b, err := xml.Marshal(act)
		if err != nil {
			glog.Errorf("encode weichat event action err %v\n", err)
		}

		w.Write(b)
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
		glog.Errorf("authValidater exchange access_token %#v\n", *tk)

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
