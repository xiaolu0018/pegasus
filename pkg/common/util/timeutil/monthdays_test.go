package timeutil

import "testing"

func TestGetAllDayFromTimePeriod(t *testing.T) {
	days, err := GetAllDayFromTimePeriod("2002-01-02", "2002-01-02")

	if err != nil {
		t.Fail()
	}

	if len(days) != 1 {
		t.Fatalf("must have one ,but have %d ", len(days))
	}

	days, err = GetAllDayFromTimePeriod("2002-01-05", "2002-02-09")
	if err != nil {
		t.Fail()
	}
	if len(days) != 36 {
		t.Fatalf("must have 35,but have %d", len(days))
	}

	days, err = GetAllDayFromTimePeriod("2004-11-01", "2005-02-25")
	if len(days) != (92 + 25) {
		t.Fatalf("must have %d,but have %d", (92 + 25), len(days))
	}
}
