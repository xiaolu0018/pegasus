package user

import (
	"errors"
	"github.com/golang/glog"
)

var ErrUnexpectedUser error = errors.New("wrong user format")

func initUser(bsonID, openID string) error {
	//return db.User().Insert(bson.M{"_id": bson.ObjectIdHex(bsonID), "openid": openID})
	u := User{
		ID:     bsonID,
		OpenID: openID,
	}
	glog.Errorln("initUser___", u.Upsert())
	return u.Upsert()
}
