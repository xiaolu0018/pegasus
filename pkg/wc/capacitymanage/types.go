package capacitymanage

import (
	"gopkg.in/mgo.v2/bson"
	"sync"
)

//分院容量管理信息
type CapacityManage struct {
	ID            bson.ObjectId `json:"-" bson:"_id,omitempty"`
	BranchID      bson.ObjectId `json:"branchid" bson:"branchid,omitempty"`
	Year          int           `json:"year"`  //年份
	Month         int           `json:"month"` //月份
	Year_Month    string        //为了查询时方便
	Mu            sync.Mutex
	DayOfCapacity map[string]int //写入某天可预约人数
	SpecialItem   map[string]int `specialitem` //特殊项目体检，key 为项目，value为项目每天可预约人数
	OffDays       []int          //写入休息的日期
}
