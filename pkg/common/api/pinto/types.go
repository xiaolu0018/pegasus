package pinto

import (
	"bjdaos/pegasus/pkg/common/types"
	"time"
)

type Appointment struct {
	ID          string   `json:"id"`
	PlanId      string   `json:"planid"` //套餐
	SaleCodes   []string `json:"sale_codes"`
	AppointTime int64    `json:"appoint_time"` //预约时间的int
	AppointDate string   `json:"appoint_date"` //预约时间的日期 yyyy-MM-dd
	OrgCode     string   `json:"org_code"`     //分院

	CardNo          string `json:"cardno"`
	CardType        string `json:"cardtype"`
	Mobile          string `json:"mobile"`
	Appointor       string `json:"appointor"` //预约人姓名
	Address         string `json:"address"`
	MerryStatus     string `json:"merrystatus"`
	Status          string `json:"status"`
	Appoint_Channel string `json:"appoint_channel"` //预约渠道
	Appointorid     string `json:"appointorid"`     //预约人id

	Sex          string `json:"sex"`
	Company      string `json:"company"`
	Group        string `json:"group"`
	Remark       string `json:"remark"`
	Operator     string `json:"operator"`
	OperateTime  int64  `json:"-"`            //创建时间
	OperateDate  string `json:"operate_date"` //yyyy-MM-dd
	OrderID      string `json:"orderid"`
	CommentID    string `json:"commentid"`
	AppointedNum int    `json:"appointednum"` //最后生产的预约号

	ReportId string `json:"reportid"` //用来记录体检报告号
	BookNo   string `json:"bookno"`
	IfSingle bool   `json:"ifsingle"` //是否散客
	IfCancel bool   `json:"ifcancel"` //是否取消预约体检
	TimeNow  time.Time
}

type ExamsAll struct {
	E *types.Examination
	Checkups []types.ExaminationCheckUp
	Sales    []types.ExaminationSale
	P *types.Person
	B *types.BookRecord
}

type ForStatistics struct {
	HosCode   []string
	OrgCode   []string //暂时不用，因为通过预约的是不分科室的
	StartDate string
	EndDate   string
	Checkups  []string
}

