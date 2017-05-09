package organization

import (
	"testing"
	"time"

	"bjdaos/pegasus/pkg/appoint/db"
	"bjdaos/pegasus/pkg/common/util/error"
	"fmt"
)

func TestConfig_Basic_Create(t *testing.T) {
	if err := db.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto"); err != nil {
		t.Fatal(err)
	}

	org := Config_Basic{
		Org_Code:     "000100131",
		Capacity:     100,
		WarnNum:      90,
		OffDays:      []string{"fdfdfdf", "健康信息"},
		AvoidNumbers: []int64{13, 3, 4, 14},
	}

	err := org.Create()
	if err == nil {
		t.Fatal(err)
	}
}

func TestConfig_Special_Create(t *testing.T) {
	if err := db.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto"); err != nil {
		t.Fatal(err)
	}

	org := &Config_Special{
		Org_Code:  time.Now().String()[:30],
		CheckupCode: time.Now().String()[:30],
		Capacity:  100,
	}

	err := org.Create()
	if err == nil {
		t.Fatal(err)
	}

	if !error.ForeignKeyConstraint(err) {
		t.Fatal(err)
	}
}

func TestORg(t *testing.T) {
	if err := db.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto"); err != nil {
		t.Fatal(err)
	}
}
