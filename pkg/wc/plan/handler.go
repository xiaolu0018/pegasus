package plan

import (
	"net/http"

	"github.com/golang/glog"

	"bjdaos/pegasus/pkg/wc/common"
	"github.com/julienschmidt/httprouter"

	httputil "github.com/1851616111/util/http"
)

func GetPlansHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rspbyte, statuscode, err := common.Go_Through_Http("GET", "/api/plans", "")
	if statuscode != 200 || err != nil {
		glog.Errorln("wc.GetPlansHandler go_through_http ", statuscode, err)
		httputil.ResponseJson(w, 400, err)
	}
	if _, err := w.Write(rspbyte); err != nil {
		httputil.ResponseJson(w, 400, err)
		return
	}
	return
}
