package vote

import (
	"github.com/julienschmidt/httprouter"
	"testing"
)

func TestNewRouter(t *testing.T) {
	var err error
	dbI, err = NewDBInterface("192.168.199.216", "5432", "postgres", "postgresql2016", "pinto")
	if err != nil {
		t.Fatal(err)
	}

	r := httprouter.New()
	AddRouter(r)
}
