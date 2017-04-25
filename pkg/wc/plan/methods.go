package plan

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//plan 保存更新
func (p *Plan) UpSert(c *mgo.Collection) error {
	if bson.IsObjectIdHex(p.Id.Hex()) {
		return c.UpdateId(p.Id, p)
	}
	p.Id = bson.NewObjectId()
	return c.Insert(p)
}

//暂时只检查是否为空
func (p Plan) Validate() error {
	if p.ImageUrl == "" || p.DetailsUrl == "" {
		return errors.New("params invalid")
	}
	return nil
}

func GetPlans(c *mgo.Collection) ([]Plan, error) {
	plans := make([]Plan, 0)
	err := c.Find(bson.M{"ifshow": false}).All(&plans)
	return plans, err
}
