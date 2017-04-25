package plan

import "gopkg.in/mgo.v2/bson"

//定义套餐结构体 Plan

type Plan struct {
	Id           bson.ObjectId `bson:"_id"`
	Name         string        `bson:"name" json:"name"`
	OrigPrice    string        `bson:"origprice" json:"origPrice"`       //团购价
	Discount     string        `bson:"discount" json:"discount"`         //折扣
	PresentPrice string        `bson:"presentprice" json:"presentPrice"` //个人价
	ImageUrl     string        `bson:"imageurl" json:"imageUrl"`
	DetailsUrl   string        `bson:"detailurl" json:"detailsUrl"`

	SpecialItems []string //一个套餐中应该含有几种检查类型

	IfShow bool `bson:"ifshow" json:"-"` //是否显示  false 显示  ture 隐藏
}
