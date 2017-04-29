package types

type BookRecord struct {
	BookNo         string `json:"bookno"`
	ExaminationNo  string `json:"examination_no"`
	Truename       string `json:"truename"`
	Sex            int    `json:"sex"`
	Bookid         string `json:"bookid"`
	Bookidtype     string `json:"bookidtype"`
	Booktimestamp  string `json:"booktimestamp"`
	BirthDay       string `json:"birthday"`
	BookorgCode    string `json:"bookorg_code"`
	AppointChannel string `json:"appoint_channel"`
	CreateTime     string `json:"createtime"`
	Telphone       string `json:"telphone"`
	BookCode       string `json:"book_code"`
}
