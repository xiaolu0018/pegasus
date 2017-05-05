package plan

import (
	"bjdaos/pegasus/pkg/wc/db"
	"fmt"
	"testing"
)

func TestPlan_UpSert(t *testing.T) {
	p := Plan{}
	p.Name = "迪安套餐"
	p.ImageUrl = "img/pack1.png"
	p.DetailsUrl = "img/pacdet3.jpg"

	err := p.UpSert(db.Plan())
	if err != nil {
		t.Fatal("no id  err ", err.Error())
	}

}

func TestGetPlans(t *testing.T) {
	_, err := GetPlans(db.Plan())
	if err != nil && err.Error() != "not found" {
		t.Fatal(err.Error())
	}
}

func TestSendHttpToGetPlans(t *testing.T) {
	sd, to := SendHttpToGetPlans()
	fmt.Println("sddto", string(sd), to)
}
