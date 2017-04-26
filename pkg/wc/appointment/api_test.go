package appointment

import (
	"192.168.199.199/bjdaos/pegasus/pkg/wc/db"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestGet(t *testing.T) {
	ap, err := Get(bson.ObjectIdHex("58f883cdc196262d2bea3963"))
	ap.UpdateStatus(db.Appointment(), ap.SpecialItem)

	fmt.Println("zheli", ap, err)
}


func TestGetListAppointmentFromApp(t *testing.T){
	GetListAppointmentFromApp("590002bbc1962674c1f5e0ad")
}