package types

type Examination struct {
	ExaminationNo   string `json:"examination_no"`
	CreateTime      string `json:"createtime"`
	UpdateTime      string `json:"updatetime"`
	Status          string `json:"status"`
	PersonCode      string `json:"person_code"`
	OrgCode         string `json:"org_code"`
	HosCode         string `json:"hos_code"`
	CheckupDate     string `json:"checkupdate"`
	CheckupHoscode  string `json:"checkup_hoscode"`
	GuidePaperState string `json:"guide_paper_state"` //微信为默认0
	ReportGrantType string `json:"report_grant_type"`
}

type ExaminationCheckUp struct {
	ExaminationNo string `json:"examination_no"`
	CheckupCode   string `json:"checkup_code"`
	CheckupStatus int    `json:"checkup_status"`
	CreateTime    string `json:"createtime"`
	HosCode       string `json:"hos_code"`
}

type ExaminationSale struct {
	ExaminationNo string  `json:"examination_no"`
	SaleCode      string  `json:"sale_code"`
	SaleStatus    string  `json:"sale_status"`
	HosCode       string  `json:"hos_code"`
	SaleSellprice float64 `json:"sale_sellprice"`
	Discount      float64 `json:"discount"`
	Curprice      float64 `json:"curprice"`
}
