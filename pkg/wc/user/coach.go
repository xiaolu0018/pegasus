package user

import (
	"errors"
	"fmt"
	"time"

	"192.168.199.199/bjdaos/pegasus/pkg/wc/token"
	"github.com/1851616111/util/weichat/manager/user"
	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
)

const WC_USER_INFO_CH_LEN = 500

var nilStruct struct{} = struct{}{}

//用户缓存数据
var IDCache Cache

func Init() error {
	IDCache = NewCache(1800)
	return IDCache.(*cache).Init()
}

type Cache interface {
	GetWorkingToken(id string) (bool, string)
	CacheSysToken(id, token string) error
	CacheWCToken(id, token string) error
	Auth(string) (bool, string)
}

func NewCache(ttlSec int64) Cache {
	c := &cache{
		tokenTTLSec:         ttlSec,
		oIDToDbID:           map[string]string{},
		oIDToToken:          map[string]string{},
		tokenToOID:          map[string]string{},
		tokenToSecs:         map[string]int64{},
		toSyncOIDs:          map[string]struct{}{},
		toSyncCh:            make(chan struct{}, 200),
		toCacheWcUserInfoCh: make(chan *wcToken, WC_USER_INFO_CH_LEN),
	}

	go c.startSync()
	return c
}

type cache struct {
	oIDToDbID  map[string]string
	toSyncOIDs map[string]struct{}
	toSyncCh   chan struct{}

	tokenTTLSec         int64
	oIDToToken          map[string]string
	tokenToOID          map[string]string //token过期删除oID
	tokenToSecs         map[string]int64  //token缓存时间秒数
	toCacheWcUserInfoCh chan *wcToken
}

func (c *cache) Init() error {
	var openIDs []string
	var err error
	if openIDs, err = user.ListUserIDs(token.TokenCtrl.GetToken()); err != nil {
		return err
	}

	if len(openIDs) == 0 {
		return nil
	}

	var users []User
	if users, err = GetUsersByOpenids(openIDs); err != nil {
		return err
	}

	c.cacheUsers(users)

	fmt.Printf("初始化微信关注用户: %d \n", len(openIDs))
	fmt.Printf("初始化微信登陆用户: %d \n", len(c.oIDToDbID))

	return nil
}

func (c *cache) GetWorkingToken(id string) (bool, string) {
	token, exist := c.oIDToToken[id]
	if !exist {
		return false, ""
	}

	sec, exist2 := c.tokenToSecs[token] //未知原因导致ｔｏｋｅｎ和时间没有同步
	if !exist2 {
		c.cleanToken(id, token)
		return false, ""
	}

	if time.Now().Unix()-sec > c.tokenTTLSec { //token过期删除
		c.cleanToken(id, token)
		return false, ""
	}

	return true, token
}

func (c *cache) CacheSysToken(id, token string) error {
	if id == "" {
		return errors.New("cache id not found")
	}
	if token == "" {
		return errors.New("cache token not found")
	}

	//openID不存在, 表示第一次登陆, 创建bsonID, 并同步.
	if _, exist := c.oIDToDbID[id]; !exist {
		c.oIDToDbID[id] = bson.NewObjectId().Hex()
		c.toSyncOIDs[id] = nilStruct
		c.toSyncCh <- nilStruct
	}

	//这个ｔｏｋｅｎ是通过GetWorkingToken获取的
	_, exist := c.oIDToToken[id]
	if !exist {
		c.oIDToToken[id] = token
		c.tokenToOID[token] = id
		c.tokenToSecs[token] = time.Now().Unix()
	}

	return nil
}

func (c *cache) CacheWCToken(id, token string) (err error) {
	if len(c.toCacheWcUserInfoCh) == WC_USER_INFO_CH_LEN {
		return errors.New("weichat user channel full")
	}

	c.toCacheWcUserInfoCh <- &wcToken{
		openID:      id,
		accessToken: token,
	}

	return
}

func (c *cache) Auth(token string) (bool, string) {
	if oid, ok := c.tokenToOID[token]; ok {
		id, ok := c.oIDToDbID[oid]
		return ok, id
	}
	return false, ""
}

func (c *cache) cacheUsers(l []User) {
	for _, user := range l {
		c.oIDToDbID[user.OpenID] = user.ID.Hex()
	}
}

func (c *cache) startSync() {
	clock := time.NewTicker(time.Second * 10)
	tokenExpirClock := time.NewTicker(time.Second * 60)
	oomCheck := time.NewTicker(time.Second * 30)
	for {
		select {
		case <-c.toSyncCh:
			c.sync()
		case <-clock.C:
			c.sync()
		case <-tokenExpirClock.C:
			c.cleanExpiredToken()
		case <-oomCheck.C:
			c.Check()
		case wcU := <-c.toCacheWcUserInfoCh:
			if err := wcU.updateUserWCInfo(); err != nil {
				glog.Errorf("cache.startSync()  sync cache weichat user err %v\n", err)
			}
		}
	}
}

func (c *cache) sync() {
	for oid := range c.toSyncOIDs {
		if err := initUser(c.oIDToDbID[oid], oid); err != nil {
			glog.Errorf("wc user cache sync user(openid=%s, bsonid=%s) err %v\n\n", oid, c.oIDToDbID[oid], err)
		} else {
			delete(c.toSyncOIDs, oid)
		}
	}
}

func (c *cache) cleanExpiredToken() {
	for token, sec := range c.tokenToSecs {
		if time.Now().Unix()-sec > c.tokenTTLSec { //token过期删除
			c.cleanByToken(token)
			return
		}
	}
}

func (c *cache) cleanToken(id, token string) {
	delete(c.oIDToToken, id)
	delete(c.tokenToOID, token)
	delete(c.tokenToSecs, token)
}

func (c *cache) cleanByToken(token string) {
	delete(c.tokenToSecs, token)
	delete(c.oIDToToken, token)
	if id, ok := c.tokenToOID[token]; ok {
		delete(c.oIDToToken, id)
	}
}

func (c *cache) Check() {
	glog.Errorf("cache reporter: openID -- bsonID: %d\n", len(c.oIDToDbID))
	glog.Errorf("cache reporter: openID -- sync: %d\n", len(c.toSyncOIDs))
	glog.Errorf("cache reporter: openID -- token: %d\n", len(c.oIDToToken))
	glog.Errorf("cache reporter: token -- openID: %d\n", len(c.tokenToOID))
	glog.Errorf("cache reporter: token -- sec: %d\n", len(c.tokenToSecs))
}
