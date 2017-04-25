package capacitymanage

import (
	"fmt"
	"time"

	"strconv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/golang/glog"

	"192.168.199.199/bjdaos/pegasus/pkg/wc/branch"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/db"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/util"
)

//容量管理初始话
func CapacityManageInit() {
	go func() {
		for {
			if time.Now().Hour() == 23 {
				go capacityManageInit()
			}
			time.Sleep(time.Hour)
		}
	}()
}

func capacityManageInit() {
	c_b, c_i := db.Branch(), db.CapacityManage()
	time := time.Now()

	yearAndMonths := afterMonths(time.Year(), int(time.Month()), 3)

	branchs, err := branch.ListBranches(c_b)
	if err != nil {
		glog.Errorln("ListBranches err", err.Error())
	}

	for _, branch := range branchs {
		for _, ym := range yearAndMonths {
			if err = InitOneMonthCpacityManage(branch.ID, ym.month, ym.year, c_b, c_i); err != nil {
				glog.Errorln("InitOneMonthCpacityManage err", err.Error())
			}
		}
	}

}

func (b *CapacityManage) Insert(c *mgo.Collection) error {
	b.Year_Month = fmt.Sprintf("%d-%2d", b.Year, b.Month)
	return c.Insert(b)
}

func FindCapacityManage(year, month int, branchid string, c *mgo.Collection) ([]CapacityManage, error) {
	yearAndMonths := afterMonths(year, month, MonthInt_) //拿到后续的MonthInt_个月时间
	star := fmt.Sprintf("%d%s", yearAndMonths[0].year, monthToString(yearAndMonths[0].month))
	end := fmt.Sprintf("%d%s", yearAndMonths[(len(yearAndMonths)-1)].year, monthToString(yearAndMonths[(len(yearAndMonths)-1)].month))
	var cms []CapacityManage
	err := c.Find(bson.M{"branchid": bson.ObjectIdHex(branchid), "year_month": bson.M{"$gte": star, "$lte": end}}).All(&cms)
	return cms, err
}

func GetCapacityManage(query interface{}) (*CapacityManage, error) {
	c := CapacityManage{}
	if err := db.CapacityManage().Find(query).One(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
func (c *CapacityManage) UpdateDayOfCapacity(day string, specialitems []string) bool {

	//修改当天的可预约的总容量
	if c.DayOfCapacity[day] == 0 {
		return false
	}
	c.DayOfCapacity[day] -= 1

	//修改当天的特殊项目的相应容量
	for _, specialitem := range specialitems {
		if value, ok := c.SpecialItem[specialitem]; ok {
			if value <= 0 {
				return false
			}
			c.SpecialItem[specialitem] -= 1
		} else {
			//在这没有发现要体检的特殊项目，正常情况是不可预约的
		}
	}

	return true
}

//初始化 某分院的具体月份的CpacityManage
func InitOneMonthCpacityManage(branchid bson.ObjectId, year, month int, c_branchinfo, c_branch *mgo.Collection) error {
	branch := branch.Branch{}
	if err := c_branch.FindId(branchid).One(&branch); err != nil {
		return err
	}

	if count, err := c_branchinfo.Find(bson.M{"branchid": branchid, "year": year, "month": month}).Count(); err != nil {
		return err
	} else {
		if count > 0 {
			return nil
		} else {
			fullinfo := initDayOfCapacity(month, year, branch.Capacity)
			branchinfo := CapacityManage{BranchID: branchid, Year: year, Month: month, DayOfCapacity: fullinfo, SpecialItem: branch.SpecialItem}
			return branchinfo.Insert(c_branchinfo)
		}
	}
}

//初始化一月中每天可预约的人数fullinfo 的map
func initDayOfCapacity(month, year, capacity int) map[string]int {
	days := util.MonthDays(month, year)
	fullinfo := make(map[string]int)
	for i := 0; i < days; i++ {
		fullinfo[strconv.Itoa(i+1)] = capacity
	}
	return fullinfo
}

type yearAndMonth struct {
	year  int
	month int
}

func monthToString(month int) string {
	if month > 9 {
		return strconv.Itoa(month)
	} else {
		return "0" + strconv.Itoa(month)
	}
}

// 得出确定月份往后n个月的结构
func afterMonths(year, month, n int) []yearAndMonth {
	yearAndMonths := make([]yearAndMonth, 0, n)
	yearandmonth := yearAndMonth{}
	for i := 0; i < n; i++ {
		if (month + i) > 12 {
			y := year + (month+i)/12

			m := (month + i) % 12
			if (month+i)%12 == 0 {
				m = 12
			}
			yearandmonth.year, yearandmonth.month = y, m
		} else {
			yearandmonth.year, yearandmonth.month = year, (month + i)
		}
		yearAndMonths = append(yearAndMonths, yearandmonth)
	}
	return yearAndMonths
}

//调整休假时间输出结构
func FilterOffDays(cms []CapacityManage) map[string]interface{} {
	result := make(map[string]interface{})
	offdays := make(map[string][]int)
	capatityman := make(map[string][]int)
	for _, cm := range cms {
		offdays[strconv.Itoa(cm.Year)+monthToString(cm.Month)] = cm.OffDays

		cap := make([]int, 0, len(cm.DayOfCapacity))
		for k, v := range cm.DayOfCapacity {
			if v == 0 {
				kint, err := strconv.Atoi(k)
				if err != nil {
					return nil
				}
				cap = append(cap, kint)
			}
		}
		capatityman[strconv.Itoa(cm.Year)+monthToString(cm.Month)] = cap
	}
	result["capatityed"] = capatityman
	result["offdays"] = offdays

	return result
}

//修改每月休息时间
func (cm *CapacityManage) UpdateOffDays(c *mgo.Collection) error {
	return c.UpdateId(cm.ID, bson.M{"$set": bson.M{"restdays": cm.OffDays}})
}
