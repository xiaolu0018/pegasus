package pinto

import (
	"testing"
)

func TestSyncController_SyncOrganizations(t *testing.T) {
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
