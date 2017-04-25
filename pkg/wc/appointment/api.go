package appointment

import (
	"fmt"

	"strconv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"

	"192.168.199.199/bjdaos/pegasus/pkg/wc/db"

	"192.168.199.199/bjdaos/pegasus/pkg/wc/branch"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/capacitymanage"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/plan"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/user"
	"github.com/golang/glog"
	"strings"
	"time"
)

func (a *Appointment) Create(c *mgo.Collection) error {
	a.ID = bson.NewObjectId()

	if bson.IsObjectIdHex(a.PlanID.Hex()) {
		plan, err := plan.Get(a.PlanID) //todo speciitem 应该从前端直接传过来，不需要在这重新
		if err != nil {
			glog.Errorln("plan err %v", a)
			return err
		}
		a.SpecialItem = plan.SpecialItems
	}

	a.CreatDate = time.Now()
	glog.Errorln("a err %v", a)
	return c.Insert(a)
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
		glog.Errorln("zheli<<<<<<<", err.Error())
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
	var err error

	b := &branch.Branch{}
	if b, err = branch.Get(a.BranchID); err != nil {
		return nil, err
	}
	au.BranchName = b.Name
	if bson.IsObjectIdHex(a.PlanID.Hex()) {
		p := &plan.Plan{}
		if p, err = plan.Get(a.PlanID); err != nil {
			return nil, err
		}
		au.Planname = p.Name
	}
	au.AppointDate = a.AppointDate
	au.Name = u.Name
	au.Mobile = u.Mobile
	au.CardID = u.CardNo
	return &au, nil
}
