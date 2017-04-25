package organization

import (
	"testing"
	"time"

	"192.168.199.199/bjdaos/pegasus/pkg/appoint/db"
	"192.168.199.199/bjdaos/pegasus/pkg/common/util/error"
	"fmt"
)

func TestConfig_Basic_Create(t *testing.T) {
	if err := db.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto"); err != nil {
		t.Fatal(err)
	}

	org := Config_Basic{
		Org_Code:     time.Now().String()[:30],
		Capacity:     100,
		WarnNum:      90,
		OffDays:      []string{"2017-02-13", "2017-02-14"},
		AvoidNumbers: []int64{13, 3, 4, 14},
	}

	err := org.Validate()
	fmt.Println("125464_____", org.AvoidNumbers, err)
	//err := org.Create()
	if err == nil {
		t.Fatal(err)
	}

	//if !error.ForeignKeyConstraint(err) {
	//	t.Fatal(err)
	//}
}

func TestConfig_Special_Create(t *testing.T) {
	if err := db.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto"); err != nil {
		t.Fatal(err)
	}

	org := &Config_Special{
		Org_Code:  time.Now().String()[:30],
		Sale_Code: time.Now().String()[:30],
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
