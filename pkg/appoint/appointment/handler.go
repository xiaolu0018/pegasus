package appointment

import (
	"time"
	"strings"
	"strconv"
	"net/http"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"

	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	httputil "github.com/1851616111/util/http"

	tm "bjdaos/pegasus/pkg/common/util/time"
)

func CreateAppointmentHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	a := Appointment{}
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		glog.Errorf("appointment.CreateAppointmentHandler Decode req params err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}

	if err := a.Validate(); err != nil {
		glog.Errorf("appointment.CreateAppointmentHandler Validate req params err %v\n", err)
		httputil.Response(rw, 400, err)
		return
	}

	if err := a.CreateAppointment(); err != nil {
		if err.Error() == ErrAppointmentString {
			httputil.Response(rw, 200, "ErrAppointmentString")
			return
		}
		glog.Errorf("orgnization.CreateHandle Create err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}

	httputil.Response(rw, 200, "ok")
}

func GetAppointmentHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	appointid := ps.ByName("appointid")
	appointment := &Appointment{}
	var err error
	if appointment, err = GetAppointment(appointid); err != nil {
		glog.Errorf("orgnization.CancelAppointmentHandler GetAppointment err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	if err := json.NewEncoder(rw).Encode(appointment); err != nil {
		httputil.Response(rw, 400, err)
		return
	}
	return

}

func ListAppointmentsHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//todo 应该区分个人查询，还是管理员查询

	rw.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	rw.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	rw.Header().Set("content-type", "application/json")

	page_no := r.FormValue("page_no")
	pageSize := r.FormValue("page_size")
	org_code := r.FormValue("org_code")
	begintimestring := r.FormValue("begintime")
	endtimestring := r.FormValue("endtime")
	search := r.FormValue("search")

	userid := ps.ByName("user")
	if strings.Contains(userid, "admin") { //是管理员操作
		userid = ""
	}

	var page_index, page_size int

	if page_no == "" {
		page_no = "0"
	}

	var err error
	if page_index, err = strconv.Atoi(page_no); err != nil {
		glog.Errorln("Orgnization ListAppointmentsHandler page_index", err.Error())
		httputil.Response(rw, 400, err)
		return
	}

	if pageSize == "" {
		pageSize = "20"
	}

	if page_size, err = strconv.Atoi(pageSize); err != nil {
		glog.Errorln("appointment ListAppointmentsHandler page_size ,err", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	var begintime, endtime time.Time
	var beginInt64, endInt64 int64
	if begintimestring != "" {
		begintime, err = time.Parse("2006-01-02", begintimestring)
		if err != nil {
			glog.Errorln("appointment ListAppointmentsHandler begintime ,err", err.Error())
			httputil.Response(rw, 400, err)
			return
		}
		beginInt64 = tm.TodayStartSec(begintime)
	} else {
		beginInt64 = 0
	}

	if endtimestring != "" {
		endtime, err = time.Parse("2006-01-02", endtimestring)
		if err != nil {
			glog.Errorln("appointment ListAppointmentsHandler endtime ,err", err.Error())
			httputil.Response(rw, 400, err)
			return
		}
		endInt64 = tm.TodayStartSec(endtime)
	} else {
		endInt64 = 0
	}

	var apps []Appointment
	var total int
	if apps, total, err = GetAppointmentList(page_index, page_size, beginInt64, endInt64, org_code, search, userid); err != nil {
		glog.Errorln("appointment ListAppointmentsHandler GetAppointmentList ,err", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	result := make(map[string]interface{})
	glog.Errorln("apps__", len(apps))
	result["total"] = total
	result["data"] = apps

	httputil.ResponseJson(rw, 200, result)
}

func ListAppointmentsForWeChatHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userid := ps.ByName("user")
	if strings.Contains(userid, "admin") { //是管理员操作
		userid = ""
	}
	var apps []Appointment
	var err error
	if apps, _, err = GetAppointmentList(0, 20, 0, 0, "", "", userid); err != nil {
		glog.Errorln("appointment ListAppointmentsHandler GetAppointmentList ,err", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	afws := GetApp_for_wechatsByAppointments(apps)
	httputil.ResponseJson(rw, 200, &afws)
	return

}

func CancelAppointmentHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	appointid := ps.ByName("appointid")
	appointment := &Appointment{}
	var err error
	if appointment, err = GetAppointment(appointid); err != nil {
		glog.Errorf("orgnization.CancelAppointmentHandler GetAppointment err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	if err = appointment.CancelAppointment(); err != nil {
		glog.Errorf("orgnization.CancelAppointmentHandler CancelAppointment err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	httputil.Response(rw, 200, "ok")
	return

}

func UpdateAppointmentHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	a := Appointment{}
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		glog.Errorf("orgnization.UpdateAppointmentHandler Decode req params err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}

	if a.Validate() != nil {
		glog.Errorf("orgnization.UpdateAppointmentHandler Validate req params err %v\n", a.Validate().Error())
		httputil.Response(rw, 400, a.Validate())
		return
	}

	if err := a.UpdateAppointment(); err != nil {
		glog.Errorf("orgnization.UpdateAppointmentHandler UpdateAppointment err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	httputil.Response(rw, 200, "ok")
	return
}

func CreateCommentHandler(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	appid := ps.ByName("appointid")
	c := Comment{}

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		glog.Errorf("appointment.CreateCommentHandler Decode req params err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}

	c.ID = bson.NewObjectIdWithTime(time.Now()).Hex()

	if err := c.Create(appid); err != nil {
		glog.Errorf("appointment.CreateCommentHandler Create err %v\n", err.Error())
		httputil.Response(rw, 400, err)
		return
	}
	httputil.Response(rw, 200, "ok")
}
