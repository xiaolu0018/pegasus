package appointment

import "testing"

func TestgetTodayOrgBookOrders(t *testing.T) {
	dbinit()
	_, err := getTodayOrgBookOrders()
	if err != nil {
		t.Fatal(err)
	}
}
