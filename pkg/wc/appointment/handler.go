package appointment

import (
	"time"

	"net/http"

	"encoding/json"

	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"

	httputil "github.com/1851616111/util/http"
	"github.com/1851616111/util/message"
	"github.com/julienschmidt/httprouter"

	"192.168.199.199/bjdaos/pegasus/pkg/appoint/cache"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/common"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/user"

	appoint_Appointment "192.168.199.199/bjdaos/pegasus/pkg/appoint/appointment"
	"fmt"
	"math/rand"
	"strconv"
)

func CreateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cm := Appointment{}
	if err := json.NewDecoder(r.Body).Decode(&cm); err != nil {
		glog.Errorln("Appointment CreateHandle Decode", err.Error())
		httputil.Response(w, 400, err)
		return
	}
	cm.ID = bson.NewObjectIdWithTime(time.Now()).Hex()
	cm.UserID = ps.ByName(common.AuthHeaderKey)
	result := make(map[string]string)
	result["appointid"] = cm.ID
	if err := message.SuccessI(w, result); err != nil {
		glog.Errorln("Appointment GetAppointmentConfirmHandler SuccessI err", err.Error())
		return
	}
	cache.Set(CACHE_TP, cm.ID, cm)
	return
}

func GetAppointmentConfirmHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	appid := ps.ByName("appointid")

	a, err := cache.Get(CACHE_TP, appid)
	if err != nil {
		glog.Errorln("Appointment GetAppointmentConfirmHandler cache.Get err ", err.Error())
		httputil.Response(w, 400, err.Error())
		return
	}
	u := &user.User{}
	if u, err = user.GetUserByid(ps.ByName(common.AuthHeaderKey)); err != nil {
		glog.Errorln("Appointment GetAppointmentConfirmHandler user.Get err", err.Error())
		httputil.Response(w, 400, err)
		return
	}

	var appoint_user *Appoint_User
	if appoint_user, err = CreatAppoint_User(a.(Appointment), *u); err != nil {
		glog.Errorln("Appointment GetAppointmentConfirmHandler appoint_user err", err.Error())
		httputil.Response(w, 400, err)
		return
	}
	if err = message.SuccessI(w, appoint_user); err != nil {
		glog.Errorln("Appointment GetAppointmentConfirmHandler SuccessI err", err.Error())
		return
	}
	return

}

func ConfirmCreatHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	appid := ps.ByName("appointid")

	app_interface, err := cache.Get(CACHE_TP, appid)
	if err != nil {
		glog.Errorln("Appointment GetAppointmentConfirmHandler cache.Get err ", err.Error())
		httputil.Response(w, 400, err.Error())
		return
	}
	app := app_interface.(Appointment)

	u := &user.User{}
	if u, err = user.GetUserByid(ps.ByName(common.AuthHeaderKey)); err != nil {
		glog.Errorln("Appointment GetAppointmentConfirmHandler user.Get err", err.Error())
		httputil.Response(w, 400, err)
		return
	}
	var a_a *appoint_Appointment.Appointment
	if a_a, err = Get_Appoint_Appointment(*u, app); err != nil {
		glog.Errorln("Appointment GetAppointmentConfirmHandler Get_Appoint_Appointment err", err.Error())
		httputil.Response(w, 400, err)
		return
	}

	rsp, err := SendToAppoint(*a_a)
	if err != nil {
		glog.Errorln("Appointment GetAppointmentConfirmHandler SendToAppoint", err)
		httputil.Response(w, 400, "ok")
	}

	httputil.Response(w, rsp.StatusCode, "ok")
}

func CancelHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	appid := ps.ByName("id")
	userid := ps.ByName(common.AuthHeaderKey)
	common.Go_Through_Http("POST", "/api/appointment/"+appid+"/cancel", userid)
	httputil.Response(w, 200, "ok")
}

func ListAppointmentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userid := ps.ByName("userid")
	rwbyte, statuscode, err := common.Go_Through_Http("GET", "/api/appointmentlist/wc", userid)
	if err != nil {
		glog.Errorln("appointment ListAppointmentHandler Go_Through_Http, err : ", err.Error())
		return
	}
	if statuscode == 200 {
		w.Write(rwbyte)
		w.WriteHeader(statuscode)
		return
	}
	glog.Errorln("appointment ListAppointmentHandler err")
	httputil.Response(w, statuscode, "request err")
}

func CreateCommentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cm := Comment{}
	if err := json.NewDecoder(r.Body).Decode(&cm); err != nil {
		glog.Errorln("appointment CreateCommentHandler Decode", err.Error())
		httputil.Response(w, 400, err)
		return
	}
	appid := ps.ByName("appointid")
	userid := ps.ByName(common.AuthHeaderKey)
	rwbyte, statuscode, err := common.Go_Through_HttpWithBody("POST", "/api/appointment/"+appid+"/comment", userid, cm)
	if err != nil {
		glog.Errorln("appointment CreateCommentHandler Go_Through_HttpWithBody, err : ", err.Error())
		return
	}
	if statuscode == 200 {
		w.Write(rwbyte)
		w.WriteHeader(statuscode)
		return
	}
	httputil.Response(w, statuscode, "request err")
}

func GetCheckNoForReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mobile := ps.ByName("mobile")
	userid := ps.ByName(common.AuthHeaderKey)
	checkno := rand.Int()
	////todo 调用短信接口发短信
	fmt.Println("mobile", mobile, checkno)

	cache.Set(CACHE_TP_Check_message, userid, 8888)
}

func GetReportByAppid(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	checkno := ps.ByName("checkno")
	userid := ps.ByName(common.AuthHeaderKey)
	cache_checkno, err := cache.Get(CACHE_TP_Check_message, userid)
	if err != nil {
		glog.Errorln("appointment GetReportByAppid cache.Get", err.Error())
		httputil.Response(w, 400, err)
		return
	}
	if checkno != strconv.Itoa(cache_checkno.(int)) {
		httputil.Response(w, 200, false)
		return
	}
	//todo 在这先查appoint服务中的预约数据
	httputil.Response(w, 200, true)
}
