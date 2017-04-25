package plan

import (
	"192.168.199.199/bjdaos/pegasus/pkg/wc/db"
	"encoding/json"
	httputil "github.com/1851616111/util/http"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func UpsertHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	p := Plan{}
	var err error
	if err = json.NewDecoder(r.Body).Decode(&p); err != nil {
		httputil.ResponseJson(w, 404, "params invalid")
		return
	}

	if err = p.Validate(); err != nil {
		httputil.ResponseJson(w, 404, err.Error())
		return
	}

	if err = p.UpSert(db.Plan()); err != nil {
		glog.Errorln("plan upsertHandler", err.Error())
		httputil.ResponseJson(w, 404, err.Error())
		return
	}
	httputil.ResponseJson(w, 200, "ok")
}

func GetPlansHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var plans []Plan
	var err error
	if plans, err = GetPlans(db.Plan()); err == nil {
		httputil.ResponseJson(w, 200, &plans)
		return
	}
	glog.Errorln("GetHandler", err.Error())
	httputil.ResponseJson(w, 400, "not found")
}
