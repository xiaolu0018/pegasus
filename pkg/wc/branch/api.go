package branch

import (
	"bjdaos/pegasus/pkg/wc/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (b *Branch) Create(c *mgo.Collection) (err error) {
	return c.Insert(b)
}

func (b *Branch) Update(c *mgo.Collection, id string) (err error) {
	b.ID = bson.ObjectIdHex(id)
	_, err = c.Upsert(bson.M{"_id": b.ID}, b)
	return
}

func ListBranches(c *mgo.Collection) ([]Branch, error) {
	l := []Branch{}
	if err := c.Find(nil).All(&l); err != nil {
		return nil, err
	}
	return l, nil
}

func Get(id bson.ObjectId) (*Branch, error) {
	b := &Branch{}
	if err := db.Branch().FindId(id).One(b); err != nil {
		return nil, err
	}
	return b, nil
}
