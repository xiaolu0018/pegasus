package plan

import (
	"bjdaos/pegasus/pkg/wc/db"
	"gopkg.in/mgo.v2/bson"
)

func Get(id bson.ObjectId) (*Plan, error) {
	p := Plan{}
	if err := db.Plan().FindId(id).One(&p); err != nil {
		return nil, err
	}
	return &p, nil
}
