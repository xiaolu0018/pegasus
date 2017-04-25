package api

import (
	"192.168.199.199/bjdaos/pegasus/pkg/pinto/db"
	"testing"
)


func TestListOrganizations(t *testing.T) {
	if err := db.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto"); err != nil {
		t.Fatal(err)
	}

	pintoDB = db.GetDB()
	if _, err := ListAllOrgs(); err != nil {
		t.Fatal(err)
	}
}

func TestGetSalesByOrgCode(t *testing.T) {
	if err := db.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto"); err != nil {
		t.Fatal(err)
	}
	pintoDB = db.GetDB()
	if _, err := GetSalesByOrgCode("000100208"); err != nil {
		t.Fatal(err)
	}
}
