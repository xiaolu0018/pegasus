package db

import (
	"gopkg.in/mgo.v2"
	"time"
)

const (
	db               = "pegasus"
	c_user           = "user"
	c_plan           = "plan"
	c_branch         = "branch"
	c_banner         = "banner"
	C_capacityManage = "capacitymanage"
	C_appointment    = "appointment"
)

var session *mgo.Session

func init() {
	var err error
	if session, err = mgo.DialWithTimeout("192.168.199.198:27017", time.Second*30); err != nil {
		panic(err)
	}
}

func User() *mgo.Collection {
	return session.DB(db).C(c_user)
}

func Banner() *mgo.Collection {
	return session.DB(db).C(c_banner)
}

//套餐数据库连接
func Plan() *mgo.Collection {
	return session.DB(db).C(c_plan)
}

func Branch() *mgo.Collection {
	return session.DB(db).C(c_branch)
}

func CapacityManage() *mgo.Collection {
	return session.DB(db).C(C_capacityManage)
}

func Appointment() *mgo.Collection {
	return session.DB(db).C(C_appointment)
}
