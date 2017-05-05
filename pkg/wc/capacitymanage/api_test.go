package capacitymanage

import (
	"bjdaos/pegasus/pkg/wc/db"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
)

func TestFindCapacityManage(t *testing.T) {
	cm, err := FindCapacityManage(time.Now().Year(), int(time.Now().Month()), "58e757f08fe64213cadb1f73", db.CapacityManage())
	dd := FilterOffDays(cm)
	fmt.Println("cm", len(cm), err, dd)
}

func TestFilterOffDays(t *testing.T) {

	db.CapacityManage().UpdateId(bson.ObjectIdHex("58ede6cf8fe64213cadc8f9c"), bson.M{"$set": bson.M{"offdays": []int{1, 15, 18}}})
}
