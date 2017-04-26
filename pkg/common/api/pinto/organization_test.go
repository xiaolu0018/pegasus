package pinto

import (
	"192.168.199.199/bjdaos/pegasus/pkg/common/util/database"
	"testing"
)

func TestListOrganizations(t *testing.T) {
	db, err := database.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := ListOrganizations(db); err != nil {
		t.Fatal(err)
	}
}