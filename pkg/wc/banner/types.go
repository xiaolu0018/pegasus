package banner

//定义Banner结构
type Banner struct {
	Pos         int    `bson:"pos"` //位置
	ImageUrl    string `bson:"imageurl"`
	RedirectUrl string `bson:"redirecturl"`
	Hide        bool   `bson:"hide"` //ture 为隐藏 false 显示
}
