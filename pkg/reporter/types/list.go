package types

type PrintInfo struct {

	//ID int	`json:"id"`	      //序号
	Ex_No     *string `json:"ex_no"`      //体检单号
	Name      *string `json:"name"`       //体检人姓名
	CardNo    *string `json:"card_no"`    //体检人身份号
	Sex       *string `json:"sex"`        //体检人性别
	Ex_CkDate *string `json:"ex_ck_date"` //体检日期

	Status     *string `json:"status"`     //体检状态
	Group      *string `json:"group"`      //团单个单
	Enterprise *string `json:"enterprise"` //体检人单位

}
type PagesNumLimitPageNo struct {
	Total_PagesNum int `json:"total_pages_num"` //总页数
	Limit          int `json:"limit"`           //分页数
	Page_No        int `json:"page_no"`         //偏移量
}

type PageAndData struct {
	Page int         `json:"page"`		 //总页数
	Data []PrintInfo `json:"data"`
}
