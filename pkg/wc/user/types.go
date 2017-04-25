package user

import (
	"errors"
	"gopkg.in/mgo.v2/bson"

	"192.168.199.199/bjdaos/pegasus/pkg/wc/common"
)

const TokenTTL = 1800
const TokenLength = 40
const TokenHeaderName = "Beartoken"

const role_admin = "admin"

var errUserConflict error = errors.New("user conflict")
var errUserSmsCodeInvalid error = errors.New("sms code invlid")

//User结构定义
type User struct {
	ID      bson.ObjectId `json:"-" bson:"_id"`
	Mobile  string        `json:"mobile" bson:"mobile"` //用户的手机号
	Name    string        `json:"name" bson:"name"`
	IDCard  string        `json:"idcard" bson:"idcard"`
	Sex     string        `json:"sex" bson:"sex,omitempty"`
	IsMarry string        `json:"ismary" bson:"ismary,omitempty"`
	IsDianziReport bool   `json:"isdianzireport"`  //是否发送电子报告
	Address common.Address
	Label   Label
	OpenID  string `json:"-" bson:"openid"`         //微信用来确认的id
	Role    string `json:"-" bson:"role,omitempty"` //暂时区分管理员和普通用户
}

//人员的相关属性
type Label struct {
	Health map[string][]string
}

//在更新label时的操作
type UserLabel struct {
	labelmap map[string][]string
}

func (u *User) IsAdmin() bool {
	return u.Role == role_admin
}
