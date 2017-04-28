package organization

const (
	TABLE_ORG             = "go_appoint_organization"
	TABLE_ORG_CON_BASIC   = "go_appoint_organization_basic_con"
	TABLE_ORG_CON_SPECIAL = "go_appoint_organization_special_con"
)

type Organization struct {
	ID         string           `json:"id"`
	Code       string           `json:"org_code"`
	Name       string           `json:"name"`
	ImageUrl   string           `json:"imageUrl"`
	DetailsUrl string           `json:"detailsUrl"`
	BasicCon   Config_Basic     `json:"basic_con,omitempty"`
	SpecialCon []Config_Special `json:"special_con,omitempty"`
}

//分院
type Config_Basic struct {
	Org_Code     string   `json:"-"`
	Capacity     int      `json:"capacity"` //每天最多可预约人数
	WarnNum      int      `json:"warnnum"`
	OffDays      []string `json:"offdays"`       //休息日
	AvoidNumbers []int64  `json:"avoid_numbers"` //不可使用的预约号
}

//组织特殊项的容量
type Config_Special struct {
	Org_Code  string `json:"org_code"`  //外键
	Sale_Code string `json:"sale_code"` //外键
	Capacity  int    `json:"capacity"`
}

type Org_WC struct {
	OrgCode    string `json:"org_code"`
	Name       string `json:"name"`
	ImageUrl   string `json:"imageUrl"`
	DetailsUrl string `json:"detailsUrl"`
}
