package appointment

import (
	"net/http"

	"encoding/json"

	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"

	httputil "github.com/1851616111/util/http"

	appoint_Appointment "192.168.199.199/bjdaos/pegasus/pkg/appoint/appointment"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/common"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/db"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/user"
	"github.com/1851616111/util/message"

	"bytes"
	"fmt"
	"time"
)

func CreateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cm := Appointment{}
	if err := json.NewDecoder(r.Body).Decode(&cm); err != nil {
		glog.Errorln("Appointment CreateHandle Decode", err.Error())
		httputil.Response(w, 400, err)
		return
	}
	cm.UserID = bson.ObjectIdHex(ps.ByName(common.AuthHeaderKey))
	if err := cm.Create(db.Appointment()); err != nil {
		glog.Errorln("Appointment CreateHandle Create", err.Error())
		httputil.Response(w, 400, err)
		return
	}
	result := make(map[string]string)
	result["appointid"] = cm.ID.Hex()
	if err := message.SuccessI(w, result); err != nil {
		glog.Errorln("Appointment GetAppointmentConfirmHandler SuccessI err", err.Error())
		return
	}
	return
}

func UpdateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	appid := ps.ByName("id")
	mapkey := make(map[string]string)
	if err := json.NewDecoder(r.Body).Decode(&mapkey); err != nil {
		glog.Errorln("Appointment SetBranchHandler Decode", err.Error())
		httputil.Response(w, 400, err)
		return
	}
	app, err := Get(bson.ObjectIdHex(appid))
	if err != nil {
		glog.Errorln("Appointment SetBranchHandler Get", err.Error())
		httputil.Response(w, 400, err)
		return
	}

	if err = app.Update(db.Appointment()); err != nil {
		glog.Errorln("Appointment SetBranchHandler SetBranch", err.Error())
		httputil.Response(w, 400, err)
		return
	}
	httputil.Response(w, 200, "ok")
	return
}

func GetAppointmentConfirmHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	appid := ps.ByName("appointid")
	u := user.User{}
	a, err := Get(bson.ObjectIdHex(appid))
	if err != nil {
		glog.Errorln("Appointment GetAppointmentConfirmHandler Get err", err.Error())
		httputil.Response(w, 400, err)
		return
	}
	if u, err = user.Get(bson.ObjectIdHex(ps.ByName(common.AuthHeaderKey))); err != nil {
		glog.Errorln("Appointment GetAppointmentConfirmHandler user.Get err", err.Error())
		httputil.Response(w, 400, err)
		return
	}

	var appoint_user *Appoint_User
	if appoint_user, err = CreatAppoint_User(*a, u); err != nil {
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
	app, err := Get(bson.ObjectIdHex(appid))
	if err != nil {
		glog.Errorln("CapacityManage ConfirmCreatHandler Get", err.Error())
		httputil.Response(w, 400, err)
		return
	}
	if err = app.UpdateStatus(db.Appointment(), app.SpecialItem); err != nil {
		glog.Errorln("CapacityManage ConfirmCreatHandler UpdateStatus", err.Error())
		httputil.Response(w, 400, err)
		return
	}

	httputil.Response(w, 200, "ok")
}

//当预约确认时走http 往appoint服务中传数据
func ConfirmAppointment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	glog.Errorln("enter ConfirmAppointment")
	a := appoint_Appointment.Appointment{
		ID:              "appoint155",
		Appointor:       "weixinfasong",
		CardNo:          "181818818184569",
		CardType:        "cardType1",
		Mobile:          "mobile1",
		MerryStatus:     "未婚",
		Status:          "预约",
		Appoint_Channel: "微信",
		AppointedNum:    0,
		PlanId:          "1",
		OrgCode:         "000101",
		AppointTime:     time.Now().Unix(),
		OperateTime:     time.Now().Unix(),
		OrderID:         "order13",
		Operator:        "operator13",
	}

	client := &http.Client{nil, nil, nil, time.Second * 10}
	var req *http.Request
	var rsp *http.Response

	var err error
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(a)
	if req, err = http.NewRequest("POST", "http://192.168.199.198:9200/api/appointment", &buf); err != nil {
		glog.Errorln("newrequest err", err)
		return
	}

	if rsp, err = client.Do(req); err != nil {
		glog.Errorln("newrequest err", err)
		return
	}
	defer rsp.Body.Close()
	err = httputil.ResponseJson(w, rsp.StatusCode, "")
	fmt.Println("ResponseJson", err)
	return

}

func CancelHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	appid := ps.ByName("id")
	if err := CancelAppoint(db.Appointment(), bson.ObjectIdHex(appid)); err != nil {
		glog.Errorln("Appointment CancelHandler CancelAppoin", err.Error())
		httputil.Response(w, 400, err)
		return
	}
	httputil.Response(w, 200, "ok")
}

func ListAppointmentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userid := ps.ByName("userid")
	apps, err := ListAppointment(bson.ObjectIdHex(userid), db.Appointment())
	if err != nil {
		glog.Errorln("CapacityManage ListAppointmentHandler ListAppointment", err.Error())
		httputil.Response(w, 400, err)
		return
	}
	if err := json.NewEncoder(w).Encode(apps); err != nil {
		httputil.Response(w, 400, err)
		return
	} else {
		httputil.Response(w, 200, nil)
	}
}
