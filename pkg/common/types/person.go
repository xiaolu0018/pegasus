package types

type Person struct {
	Sex            string `json:"sex"`
	CardNo         string `json:"card_no"`
	IsMarry        string `json:"is_marry"`
	Name           string `json:"name"`
	CellPhone      string `json:"cellphone"`
	CreateTime     string `json:"createtime"`
	PersonCode     string `json:"person_code"`
	IdcardTypeCode string `json:"idcard_type_code"`
	HosCode        string `json:"hos_code"`
}
