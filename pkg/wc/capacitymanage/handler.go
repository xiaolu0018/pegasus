package capacitymanage

import (
	"192.168.199.199/bjdaos/pegasus/pkg/wc/common"
	httputil "github.com/1851616111/util/http"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GetOffDaysHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	//cms, err := FindCapacityManage(time.Now().Year(), int(time.Now().Month()), id, db.CapacityManage())
	//if err != nil {
	//	glog.Errorln("GetOffDaysHandle" + err.Error())
	//	httputil.Response(w, 400, err)
	//	return
	//}
	//offdays := FilterOffDays(cms)
	//json.NewEncoder(w).Encode(offdays)
	//return
	rspbyte, statuscode, err := common.Go_Through_Http("/api/offday/" + id)
	if statuscode != 200 {
		httputil.ResponseJson(w, 400, err)
	}
	if _, err := w.Write(rspbyte); err != nil {
		httputil.ResponseJson(w, 400, err)
		return
	}

	return
}
