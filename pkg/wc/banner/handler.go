package banner

import (
	"net/http"
	"encoding/json"

	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"

	"192.168.199.199/bjdaos/pegasus/pkg/wc/db"
	httputil "github.com/1851616111/util/http"
)

func UpsertHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	b := Banner{}
	var err error
	if err = json.NewDecoder(r.Body).Decode(&b); err != nil {
		httputil.ResponseJson(w, 404, "params invalid")
		return
	}
	if err = b.Validate(); err != nil {
		httputil.ResponseJson(w, 400, err.Error())
		return
	}

	if err = b.CreateOrUpdate(db.Banner()); err != nil {
		httputil.ResponseJson(w, 400, err.Error())
		return
	}
	httputil.ResponseJson(w, 200, "ok")
}

func GetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var banners []Banner
	var err error
	if banners, err = GetShowBanners(db.Banner()); err == nil {
		httputil.ResponseJson(w, 200, &banners)
		return
	}
	glog.Errorln("GetHandler", err.Error())
	httputil.ResponseJson(w, 400, "not found")
}
