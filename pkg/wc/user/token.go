package user

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"

	"bjdaos/pegasus/pkg/wc/db"
	"github.com/golang/glog"
)

var ErrUnexpectedUser error = errors.New("wrong user format")

func GetUser(token string) (*User, error) {
	u := User{}
	if err := db.User().Find(bson.M{"token": token}).One(&u); err != nil {
		return nil, err
	}

	if u.Mobile == "" {
		return nil, ErrUnexpectedUser
	}

	return &u, nil
}

func TryRefreshToken(openID, token string) error {
	//var err error
	//if u, err = getUserByOpenID(openID); err != nil {
	//	if err == mgo.ErrNotFound {
	//		return updateTokenByOpenID(openID, token)
	//	}
	//	return err
	//}
	//
	//if !u.tokenExist() || u.tokenExpired() {
	//	return updateTokenByOpenID(openID, token)
	//}

	return nil
}

func updateTokenByOpenID(openID, token string) (err error) {
	_, err = db.User().Upsert(bson.M{"openid": openID},
		bson.M{"$set": bson.M{"token": token, "tokenSec": time.Now().Unix()}})
	return err
}

func getUserByOpenID(openID string) (*User, error) {
	u := User{}
	if err := db.User().Find(bson.M{"openid": openID}).One(&u); err != nil {
		return nil, err
	}

	return &u, nil
}

func initUser(bsonID, openID string) error {
	//return db.User().Insert(bson.M{"_id": bson.ObjectIdHex(bsonID), "openid": openID})
	u := User{
		ID:     bsonID,
		OpenID: openID,
	}
	glog.Errorln("initUser___", u.Upsert())
	return u.Upsert()
}
