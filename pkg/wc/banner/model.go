package banner

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (b *Banner) CreateOrUpdate(c *mgo.Collection) (err error) {
	//对于同一个pos进行upsert操作
	if _, err = c.Upsert(bson.M{"pos": b.Pos}, b); err != nil {
		return fmt.Errorf("banner create insert err ", err.Error())
	}
	return
}

//通过已知条件来查询相关的banner
func FindBanners(c *mgo.Collection, query bson.M) (banners []Banner, err error) {
	banners = make([]Banner, 0, 5)
	if err = c.Find(query).Sort("pos").All(&banners); err != nil {
		return nil, fmt.Errorf("banner FindBanners", err.Error())
	}
	return
}

//通过已知条件来查询相关的banner
func GetShowBanners(c *mgo.Collection) (banners []Banner, err error) {
	banners = make([]Banner, 0, 5)
	if err = c.Find(bson.M{"hide": false}).Sort("pos").All(&banners); err != nil {
		return nil, fmt.Errorf("banner FindBanners", err.Error())
	}
	return
}
