package banner

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"bjdaos/pegasus/pkg/wc/common"
	httputil "github.com/1851616111/util/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	////var banners []Banner
	////var err error
	////if banners, err = GetShowBanners(db.Banner()); err == nil {
	////	httputil.ResponseJson(w, 200, &banners)
	////	return
	////}
	////glog.Errorln("GetHandler", err.Error())
	////httputil.ResponseJson(w, 400, "not found")
	//http.Redirect(w, r, common.AppointServe+"/api/banners", 302)
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
