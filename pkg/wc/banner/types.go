package banner

//定义Banner结构
type Banner struct {
	Pos         int    `json:"pos" bson:"pos"` //位置
	ImageUrl    string `json:"imageUrl" bson:"imageurl"`
	RedirectUrl string `json:"redirectUrl" bson:"redirecturl"`
	Hide        bool   `json:"" bson:"hide"` //ture 为隐藏 false 显示
}
