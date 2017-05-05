package plan

import (
	"bjdaos/pegasus/pkg/wc/common"
	"bjdaos/pegasus/pkg/wc/db"
	"encoding/json"
	"fmt"
	httputil "github.com/1851616111/util/http"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"time"
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

	rspbyte, statuscode := SendHttpToGetPlans()
	if statuscode != 200 {
		httputil.ResponseJson(w, 400, nil)
	}
	if _, err := w.Write(rspbyte); err != nil {
		httputil.ResponseJson(w, 400, err)
		return
	}

	return
}

func SendHttpToGetPlans() ([]byte, int) {
	client := &http.Client{nil, nil, nil, time.Second * 10}
	var req *http.Request
	var rsp *http.Response

	var err error
	if req, err = http.NewRequest("GET", common.AppointServe+"/api/plans", nil); err != nil {
		glog.Errorln("newrequest err", err)
		return nil, 400
	}

	if rsp, err = client.Do(req); err != nil {
		glog.Errorln("newrequest err", err)
		return nil, 400
	}

	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println("err", err.Error())
		return nil, 400
	}
	defer rsp.Body.Close()
	return buf, rsp.StatusCode
}
