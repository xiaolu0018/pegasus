package user

import (
	"bjdaos/pegasus/pkg/wc/common"
)

const (
	TokenTTL        = 1800
	TokenLength     = 40
	TokenHeaderName = "Beartoken"
	role_admin      = "admin"
	TABLE_USER      = "go_weixin_user"
	TABLE_HEALTH    = "go_weixin_user_health"
)

//User结构定义
type User struct {
	ID             string         `json:"-"`
	Mobile         string         `json:"mobile"` //用户的手机号
	Name           string         `json:"name" `
	CardNo         string         `json:"idcard" `
	CardType       string         `json:"cardtype"`
	Sex            string         `json:"sex" bson:"sex,omitempty"`
	IsMarry        string         `json:"ismarry" bson:"ismarry"`
	IsDianziReport bool           `json:"isdianzireport"` //是否发送电子报告
	Address        common.Address `json:"address" bson:"address"`
	OpenID         string         `json:"-" bson:"openid"`         //微信用来确认的id
	Role           string         `json:"-" bson:"role,omitempty"` //暂时区分管理员和普通用户
	WC_Info        WCUserInfo     //微信用户数据来自微信
}

//人员的相关属性
type Health struct {
	Id                     string   `json:"id"`
	Past_history           []string `json:"past_history"`           //既往史
	Family_medical_history []string `json:"family_medical_history"` //家族病史
	Exam_frequency         []string `json:"exam_frequency"`         // 体检情况
	Past_exam_exception    []string `json:"past_exam_exception"`    //既往体检异常情况
	Psychological_pressure []string `json:"psychological_pressure"` //心理压力
	Food_habits            []string `json:"food_habits"`            //饮食习惯
	Eating_habits          []string `json:"eating_habits"`          //进食习惯
	Drink_habits           []string `json:"drink_habits"`           //饮酒习惯
	Smoke_history          []string `json:"smoke_history"`          //吸烟史
}


func (u *User) IsAdmin() bool {
	return u.Role == role_admin
}
