package plan

//定义套餐结构体 Plan

type Plan struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	OrigPrice    string `json:"origPrice"`    //团购价
	Discount     string ` json:"discount"`    //折扣
	PresentPrice string `json:"presentPrice"` //个人价
	ImageUrl     string ` json:"imageurl"`
	DetailsUrl   string ` json:"detailsurl"`

	SpecialItems []string //一个套餐中应该含有几种检查类型

	IfShow bool `bson:"ifshow" json:"-"` //是否显示  false 显示  ture 隐藏
}
