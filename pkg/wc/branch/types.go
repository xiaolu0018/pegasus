package branch

import (
	"bjdaos/pegasus/pkg/wc/common"
	"gopkg.in/mgo.v2/bson"
)

//分院
type Branch struct {
	ID          bson.ObjectId  `json:"id" bson:"_id,omitempty"`
	Name        string         `json:"name"`
	Desc        string         `json:"desc"`
	Tel         string         `json:"tel"`
	ImageUrl    string         `bson:"imageurl" json:"imageUrl"`
	DetailsUrl  string         `bson:"detailurl" json:"detailsUrl"`
	Capacity    int            `json:"capacity"` //每天最多可预约人数
	SpecialItem map[string]int `specialitem`     //特殊项目体检，key 为项目，value为项目每天可预约人数
	Address     common.Address `json:"address"`
}
