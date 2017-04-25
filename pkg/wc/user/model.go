package user

import (
	"192.168.199.199/bjdaos/pegasus/pkg/wc/db"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//User 保存和修改基本信息

func (u *User) UpsertBasicInfo(c *mgo.Collection) (err error) {
	err = c.UpdateId(u.ID, bson.M{
		"$set": bson.M{
			"name":    u.Name,
			"idcard":  u.IDCard,
			"mobile":  u.Mobile,
			"adress":  u.Address,
			"sex":     u.Sex,
			"ismarry": u.IsMarry,
		},
	})
	return
}

//user 保存label信息
func UpdateLabel(c *mgo.Collection, userid bson.ObjectId, healthmap map[string][]string) error {
	return c.UpdateId(userid, bson.M{"$set": bson.M{"label.health": healthmap}})
}

//user 用户查询
func Get(userid bson.ObjectId) (User, error) {
	u := User{}
	err := db.User().FindId(userid).One(&u)
	if err != nil {
		err = fmt.Errorf("GetUser", err.Error())
	}
	return u, err
}

func listUsersByOpenIDs(c *mgo.Collection, openIDs []string) ([]User, error) {
	l := []User{}
	if err := c.Find(bson.M{"openid": bson.M{"$in": openIDs}}).All(&l); err != nil {
		return nil, err
	}
	return l, nil
}
