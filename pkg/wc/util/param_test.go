package util

import (
	"github.com/julienschmidt/httprouter"
	"testing"
)

func TestAddParam(t *testing.T) {
	ps := httprouter.Params{}
	ps = AddParam(ps, "10", "20")
	if ps.ByName("10") != "20" {
		t.Fail()
	}
}

func TestMonthDays(t *testing.T) {
	i := MonthDays(2, 2020)
	if i != 29 {
		t.Fatal("时间有错")
	}
}
