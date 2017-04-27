package pinto

import (
	"192.168.199.199/bjdaos/pegasus/pkg/appoint/db"
	"fmt"
	"testing"
)

func dbinit() {
	if err := db.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto"); err != nil {
		fmt.Println("dbinit", err)
	}
}
func TestSyncController_SyncOrganizations(t *testing.T) {
	dbinit()
	c := Config{
		User:     "postgres",
		Password: "postgres190@",
		IP:       "10.1.0.190",
		Port:     "5432",
		Database: "pinto",
	}

	if err := c.sync(); err != nil {
		t.Fatal(err)
	}
}
