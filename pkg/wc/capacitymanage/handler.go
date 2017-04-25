package capacitymanage

import (
	"192.168.199.199/bjdaos/pegasus/pkg/wc/db"
	"encoding/json"
	httputil "github.com/1851616111/util/http"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func GetOffDaysHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	cms, err := FindCapacityManage(time.Now().Year(), int(time.Now().Month()), id, db.CapacityManage())
	if err != nil {
		glog.Errorln("GetOffDaysHandle" + err.Error())
		httputil.Response(w, 400, err)
		return
	}
	offdays := FilterOffDays(cms)
	json.NewEncoder(w).Encode(offdays)
	return
}
