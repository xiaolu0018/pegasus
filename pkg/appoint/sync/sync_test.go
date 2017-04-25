package organization

import (
	"testing"
)

func TestSyncController_SyncOrganizations(t *testing.T) {
	c := SyncController{
		User:     "postgres",
		Password: "postgres190@",
		IP:       "10.1.0.190",
		Port:     "5432",
		Database: "pinto",
	}

	if err := c.syncOrganizations(); err != nil {
		t.Fatal(err)
	}
}
