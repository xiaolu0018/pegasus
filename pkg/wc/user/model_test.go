package user

import (
	"192.168.199.199/bjdaos/pegasus/pkg/wc/common"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/db"
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	var u User
	u.Mobile = "1234567"
	u.Password = "123"

	if err := u.Create(db.User()); err != nil {
		fmt.Println("testUser_err", err.Error())
	}
}

func TestCheckUserForLogin(t *testing.T) {
	mobile, pw := "55555", "123"
	if ok, s := CheckUserForLogin(db.User(), mobile, pw); ok {
		fmt.Print("ok", s)
	} else {
		fmt.Println("no ok", s)
	}
}

func TestUser_UpsertInfo(t *testing.T) {
	u := User{}
	u.OpenID = "777"
	u.IDCard = "610481199103092214"
	u.Address = common.Address{Province: "shanxi", City: "xianyang", District: "xingping", Details: "xiwenfang"}
	err := u.CreateValidate()
	fmt.Println(err)
	err = u.UpsertInfo(db.User())
	fmt.Println(err)
}

func TestGetUserByopenid(t *testing.T) {
	u, err := GetUserByMobile(db.User(), "777")
	fmt.Println(u, err)
}

func TestUser_UpdateLabel(t *testing.T) {
	labelhealth := make(map[string][]string)
	labelhealth["bingshi"] = []string{"gaoxueya", "gaoxuezhi"}
	labelhealth["jiazushi"] = []string{"guanxinbing", "naogeng"}
	err := UpdateLabel(db.User(), 777, labelhealth)
	fmt.Println(err)
}
