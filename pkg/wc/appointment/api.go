package appointment

import (
	"fmt"
	"time"

	//"bytes"
	http "net/http"
	//
	//"encoding/json"
	//"github.com/golang/glog"

	"bjdaos/pegasus/pkg/appoint/appointment"
	"bjdaos/pegasus/pkg/wc/user"
	http185 "github.com/1851616111/util/http"
)

func CreatAppoint_User(a Appointment, u user.User) (*Appoint_User, error) {
	var au Appoint_User
	au.BranchName = a.BranchName
	au.Planname = a.PlanName
	au.AppointDate = a.AppointDate
	au.Name = u.Name
	au.Mobile = u.Mobile
	au.CardID = u.CardNo
	return &au, nil
}

func SendToAppoint(a appointment.Appointment, ip string) (*http.Response, error) {
	//client := &http.Client{nil, nil, nil, time.Second * 10}
	//var req *http.Request
	//var rsp *http.Response
	//
	//var err error
	//var buf bytes.Buffer
	//json.NewEncoder(&buf).Encode(a)
	//if req, err = http.NewRequest("POST", "http://192.168.199.198:9200/api/appointment", &buf); err != nil {
	//	glog.Errorln("newrequest err", err)
	//	return nil, err
	//}
	//
	//if rsp, err = client.Do(req); err != nil {
	//	glog.Errorln("newrequest err", err)
	//	return nil, err
	//}

	return http185.Send(&http185.HttpSpec{
		URL:         ip + "/api/sales/checkups",
		Method:      "GET",
		ContentType: http185.ContentType_JSON,
		BodyParams:  http185.NewBody().Add("salecodes", &a),
	})

}

func Get_Appoint_Appointment(u user.User, a Appointment) (*appointment.Appointment, error) {
	var appointtimeint int64
	if appointtime, err := time.Parse("2006-01-02", a.AppointDate); err != nil {
		return nil, err
	} else {
		appointtimeint = appointtime.Unix()
	}
	address := fmt.Sprintf("%s-%s-%s-%s", u.Address.Province, u.Address.City, u.Address.District, u.Address.Details)
	a_a := appointment.Appointment{
		ID:          a.ID,
		PlanId:      a.PlanID,
		AppointTime: appointtimeint,
		OrgCode:     a.BranchID,

		CardNo:          u.CardNo,
		CardType:        u.CardType,
		Mobile:          u.Mobile,
		Appointor:       u.Name,
		Appointorid:     u.ID,
		Address:         address,
		MerryStatus:     u.IsMarry,
		Status:          "预约成功",
		Appoint_Channel: "微信",
		Company:         "",
		Group:           "",
		Remark:          "",
		Operator:        "微信用户",
		OperateTime:     time.Now().Unix(),
		OrderID:         "",
		CommentID:       "",
		AppointedNum:    0,
		IfSingle:        true,
		IfCancel:        false,
	}
	return &a_a, nil
}

//func GetListAppointmentFromApp(userid string) *http.Response {
//	client := &http.Client{nil, nil, nil, time.Second * 10}
//	var req *http.Request
//	var rsp *http.Response
//
//	var err error
//	if req, err = http.NewRequest("GET", "http://192.168.199.198:9200/api/appointmenlist", nil); err != nil {
//		glog.Errorln("newrequest err", err)
//		return nil
//	}
//	req.SetBasicAuth(userid, appoint.PASSWORD)
//	if rsp, err = client.Do(req); err != nil {
//		glog.Errorln("newrequest err", err)
//		return nil
//	}
//	defer rsp.Body.Close()
//	result := []map[string]interface{}{}
//	buf, err := ioutil.ReadAll(rsp.Body)
//	if err != nil {
//		fmt.Println("err", err.Error())
//		return nil
//	}
//	return rsp
//}
