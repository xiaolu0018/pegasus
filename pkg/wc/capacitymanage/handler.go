package capacitymanage

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"time"
	"github.com/golang/glog"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/db"
	httputil "github.com/1851616111/util/http"
)


func GetOffDaysHandle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	cms, err := FindCapacityManage(time.Now().Year(), int(time.Now().Month()), id, db.CapacityManage())
	if err != nil{
		glog.Errorln("GetOffDaysHandle"+err.Error())
		httputil.Response(w, 400, err)
		return
	}
	offdays := FilterOffDays(cms)
	httputil.Response(w, 200, offdays)
}
