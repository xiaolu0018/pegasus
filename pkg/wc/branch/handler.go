package branch

import (
	"net/http"
	"bjdaos/pegasus/pkg/wc/common"
	httputil "github.com/1851616111/util/http"
	"github.com/julienschmidt/httprouter"
)
// curl 192.168.199.168:9000/api/manager/branches
func ListHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rspbyte, statuscode, err := common.Go_Through_Http("GET", "/api/organizations/wc", "")
	if statuscode != 200 {
		httputil.ResponseJson(w, 400, err)
	}
	if _, err := w.Write(rspbyte); err != nil {
		httputil.ResponseJson(w, 400, err)
		return
	}

	return
}
