package pinto

import (
	"bjdaos/pegasus/pkg/common/util/database"
	"testing"
)

func TestListCheckups(t *testing.T) {
	db, err := database.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := ListCheckups(db); err != nil {
		t.Fatal(err)
	}
}

func TestGetCheckupBySaleCode(t *testing.T) {
	db, err := database.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := GetCheckupBySaleCode(db, "00000902170036"); err != nil {
		t.Fatal(err)
	}
}
