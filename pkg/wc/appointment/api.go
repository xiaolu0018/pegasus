package appointment

import (
	"fmt"

	"strconv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"

	"192.168.199.199/bjdaos/pegasus/pkg/wc/db"

	"192.168.199.199/bjdaos/pegasus/pkg/appoint"
	"192.168.199.199/bjdaos/pegasus/pkg/appoint/appointment"
	//"192.168.199.199/bjdaos/pegasus/pkg/wc/branch"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/capacitymanage"
	//"192.168.199.199/bjdaos/pegasus/pkg/wc/plan"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/user"
	"bytes"
	"encoding/json"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func (a *Appointment) Create(c *mgo.Collection) error {
	//a.ID = bson.NewObjectId().Hex()
	//
	//if bson.IsObjectIdHex(a.PlanID) {
	//	plan, err := plan.Get(a.PlanID) //todo speciitem 应该从前端直接传过来，不需要在这重新
	//	if err != nil {
	//		glog.Errorln("plan err %v", a)
	//		return err
	//	}
	//	a.SpecialItem = plan.SpecialItems
	//}
	//
	//a.CreatDate = time.Now()
	//glog.Errorln("a err %v", a)
	return nil
}

func (a *Appointment) Update(c *mgo.Collection) error {
	return c.UpdateId(a.ID, bson.M{"$set": bson.M{"appointdate": a.AppointDate, "branchid": a.BranchID}})
}

func Get(id bson.ObjectId) (*Appointment, error) {
	a := Appointment{}
	if err := db.Appointment().FindId(id).One(&a); err != nil {
		return nil, err
	}
	return &a, nil
}

func (a *Appointment) UpdateStatus(c *mgo.Collection, speitem []string) error { //需要同步更细该分院的当天可预约人数
	datestring := strings.Split(a.AppointDate, "-")
	fmt.Println(datestring)
	year, _ := strconv.Atoi(datestring[0])
	month, _ := strconv.Atoi(datestring[1])
	cm, err := capacitymanage.GetCapacityManage(bson.M{"branchid": a.BranchID, "year": year, "month": month})
	if err != nil {
		return err
	}
	if ok := cm.UpdateDayOfCapacity(datestring[2], a.SpecialItem); !ok {
		return fmt.Errorf("no capacity")
	}
	fmt.Println("zheli,,,")
	runner := txn.NewRunner(c)
	Ops := []txn.Op{
		{
			C:      db.C_capacityManage,
			Id:     cm.ID,
			Update: bson.M{"$set": bson.M{"dayofcapacity": cm.DayOfCapacity, "specialitem": cm.SpecialItem}},
		},
		{
			C:      db.C_appointment,
			Id:     a.ID,
			Update: bson.M{"$set": bson.M{"status": true}},
		},
	}
	fmt.Println("zheli,,fmtfmtfmtfmt,")
	return runner.Run(Ops, "", nil)
}

func CancelAppoint(c *mgo.Collection, id bson.ObjectId) error {
	return c.UpdateId(id, bson.M{"$set": bson.M{"cancel": true}})
}

//得到预约列表
func ListAppointment(userid bson.ObjectId, c *mgo.Collection) ([]Appointment, error) {
	apps := []Appointment{}
	if err := c.Find(bson.M{"userid": userid, "status": true}).All(&apps); err != nil {
		return nil, err
	}
	return apps, nil
}

func CreatAppoint_User(a Appointment, u user.User) (*Appoint_User, error) {
	var au Appoint_User
	//var err error

	//b := &branch.Branch{}
	//if b, err = branch.Get(a.BranchID); err != nil {
	//	return nil, err
	//}
	//au.BranchName = b.Name
	//if bson.IsObjectIdHex(a.PlanID) {
	//	p := &plan.Plan{}
	//	if p, err = plan.Get(a.PlanID); err != nil {
	//		return nil, err
	//	}
	//	au.Planname = p.Name
	//}
	au.BranchName = a.BranchName
	au.Planname = a.PlanName
	au.AppointDate = a.AppointDate
	au.Name = u.Name
	au.Mobile = u.Mobile
	au.CardID = u.CardNo
	return &au, nil
}

func SendToAppoint(a appointment.Appointment) (*http.Response, error) {
	client := &http.Client{nil, nil, nil, time.Second * 10}
	var req *http.Request
	var rsp *http.Response

	var err error
	var buf bytes.Buffer
	a.OrgCode = "000100102"
	a.PlanId = "2"
	json.NewEncoder(&buf).Encode(a)
	if req, err = http.NewRequest("POST", "http://192.168.199.198:9200/api/appointment", &buf); err != nil {
		glog.Errorln("newrequest err", err)
		return nil, err
	}

	if rsp, err = client.Do(req); err != nil {
		glog.Errorln("newrequest err", err)
		return nil, err
	}
	defer rsp.Body.Close()
	return rsp, nil
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
		Status:          "预约",
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

func GetListAppointmentFromApp(userid string) *http.Response {
	client := &http.Client{nil, nil, nil, time.Second * 10}
	var req *http.Request
	var rsp *http.Response

	var err error
	if req, err = http.NewRequest("GET", "http://192.168.199.198:9200/api/appointmenlist", nil); err != nil {
		glog.Errorln("newrequest err", err)
		return nil
	}
	req.SetBasicAuth(userid, appoint.PASSWORD)
	if rsp, err = client.Do(req); err != nil {
		glog.Errorln("newrequest err", err)
		return nil
	}
	defer rsp.Body.Close()
	result := []map[string]interface{}{}
	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println("err", err.Error())
		return nil
	}

	fmt.Println("bufff", string(buf))
	//if err := json.NewDecoder(buf).Decode(&result); err != nil {
	//	fmt.Println("new err", err)
	//}

	fmt.Println("res", result, rsp.StatusCode)
	return rsp
}
