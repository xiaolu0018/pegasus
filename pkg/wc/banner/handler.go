package banner

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"bjdaos/pegasus/pkg/wc/common"
	httputil "github.com/1851616111/util/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rspbyte, statuscode, err := common.Go_Through_Http("GET", "/api/banners", "")
	if statuscode != 200 {
		httputil.ResponseJson(w, 400, err)
	}
	if _, err := w.Write(rspbyte); err != nil {
		httputil.ResponseJson(w, 400, err)
		return
	}
	return
}
