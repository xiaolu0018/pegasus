package appointment

const TABLE_APPOINTMENT = "go_appoint_appointment"
const TABLE_BANNER = "go_appoint_banner"

const VALIDATE_CHANNEL_WC = "微信"
const VALIDATE_CHANNEL_400 = "400"
const VALIDATE_CARD_TYPE_ID = "身份证"
const VALIDATE_CARD_TYPE_PASSPORT = "护照"
const VALIDATE_CARD_TYPE_OFFICER = "军官证"
const VALIDATE_CARD_TYPE_POLICE = "警察证"
const VALIDATE_CARD_TYPE_OTHER = "其他"
const VALIDATE_MERRY_NO = "未婚"
const VALIDATE_MERRY_YES = "已婚"

var ErrAppointmentString = "Can't make an appointment"

type Appointment struct {
	ID          string   `json:"id"`
	PlanId      string   `json:"planid"` //套餐
	SaleCodes   []string `json:"sale_codes"`
	AppointTime int64    `json:"appointtime"`
	OrgCode     string   `json:"org_code"` //分院

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
	OperateTime  int64  `json:"operatetime"` //创建时间
	OrderID      string `json:"orderid"`
	CommentID    string `json:"commentid"`
	AppointedNum int    `json:"appointednum"` //最后生产的预约号

	ReportId string `json:"reportid"` //用来记录体检报告号

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
	Conclusion  string
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
	SaleCode string //todo 这个应该为checkup
	Used     int
	OrgCode  string
}

//套餐
var TABLE_PALN = "go_appoint_plan"

type Plan struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	AvatarImg string   `json:"imageurl"` //todo 暂时为了保持和微信一致
	DetailImg string   `json:"detailsurl"`
	SaleCodes []string //一个套餐中应该含有几种检查类型
	IfShow    bool     //是否显示  false 显示  ture 隐藏
}

//定义Banner结构
type Banner struct {
	Pos         int    `json:"pos" bson:"pos"` //位置
	ImageUrl    string `json:"imageUrl" bson:"imageurl"`
	RedirectUrl string `json:"redirectUrl" bson:"redirecturl"`
	IfShow      bool   //是否显示  false 显示  ture 隐藏
}

type App_For_WeChat struct {
	Appid        string `json:"appid"`
	Name         string `json:"name"`
	PlanId       string `json:"planid"`
	Org_code     string `json:"org_code"`
	Org_Name     string `json:"org_name"`
	Serve_Mobile string `json:"serve_mobile"`
	AppointDate  string `json:"appointdate"`
	OperateTime  string `json:"operatetime"`
	Reportid     string `json:"reportid"`
	Status       string `json:"status"`
}
