package appointment

var TABLE_Appointment = "go_appoint_appointment"
var ErrAppointmentString = "Can't make an appointment"

type Appointment struct {
	ID          string `json:"id"`
	PlanId      string `json:"planid"` //套餐
	AppointTime int64  `json:"appointtime"`
	OrgCode     string `json:"orgcode"` //分院

	CardNo               string `json:"cardno"`
	CardType             string `json:"cardtype"`
	Mobile               string `json:"mobile"`
	Appointor            string `json:"appointor"`
	MerryStatus          string `json:"merrystatus"`
	Status               string `json:"status"`
	Appoint_Channel      string `json:"appoint_channel"`      //预约渠道
	Channel_Appointor_ID string `json:"channel_appointor_id"` //不同渠道预约人的id

	Company      string `json:"company"`
	Group        string `json:"group"`
	Remark       string `json:"remark"`
	Operator     string `json:"operator"`
	OperateTime  int64  `json:"operatetime"` //创建时间
	OrderID      string `json:"orderid"`
	CommentID    string `json:"commentid"`
	AppointedNum int    `json:"appointednum"` //最后生产的预约号

	IfSingle bool `json:"ifsingle"` //是否散客
	IfCancel bool `json:"ifcancel"` //是否取消预约体检
}

var TABLE_Appoint_Comment = "go_appoint_comment"

type Comment struct { //预约评价
	ID          string
	Environment float32 //环境
	Attitude    float32 //态度
	Breakfast   float32 //早餐
	Details     string  //评价内容
}

//分院的某天已预约人数
var TABLE_CapacityRecords = "go_appoint_capacity_records"

type ManagerCapacity struct {
	Date    string
	OrgCode string
	Used    int
}

//分院的特殊项目的某天已预约人数
var TABLE_SaleRecords = "go_appoint_sale_records"

type ManagerItem struct {
	Date     string
	SaleCode string
	Used     int
	OrgCode  string
}

//套餐
var TABLE_PALN = "go_appoint_plan"

type Plan struct {
	ID        string
	Name      string
	AvatarImg string
	DetailImg string
	Sales     []string //一个套餐中应该含有几种检查类型
	IfShow    bool     //是否显示  false 显示  ture 隐藏
}
