package appointment

import (
	"time"
)

type Appointment struct {
	ID                string    `bson:"_id,omitempty"`
	UserID            string    `json:"userid"`
	CreatDate         time.Time `json:"CreatDate"`                      //预约创建时间按
	AppointDate       string    `json:"appointdate" bson:"appointdate"` //预约体检时间
	BranchID          string    `json:"branchid" bson:"branchid"`
	BranchName        string    `json:"branchname"`
	PlanName          string    `json:"planname"`
	PlanID            string    `bson:"planid,omitempty"`
	SpecialItem       []string  `bson:"SpecialItem"`
	AppointmentStatus string    ``                     //预约状态， 如 预约，已体检，爽约
	Form              string    `json:"form"`          //预约来源
	Oid               string    `bson:"oid,omitempty"` //是用来修改预约的，如没点最后的确认则会将这个删
	Appraise
	Status bool //是否已经确认预约   指的是在创建预约是是否点击过最后一步确认按钮
	Canel  bool //是否取消预约体检
}

type Appraise struct { //预约评价
	Type1 string
	Type2 string
	Type3 string
}

//为了预约确认时的展示
type Appoint_User struct {
	Name        string `json:"name"`
	BranchName  string `json:"branchname"`
	Planname    string `json:"planname"`
	AppointDate string `json:"appointdate"`
	CardID      string `json:"cardid"`
	Mobile      string `json:"mobile"`
}
