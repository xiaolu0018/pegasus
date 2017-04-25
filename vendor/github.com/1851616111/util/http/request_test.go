package http

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test_NewParams(t *testing.T) {
	NewParams()
}

func Test_SortParams(t *testing.T) {
	p1 := NewParams()
	p1.Add("bbb", "44444")
	p1.Set("bbb", "222")
	p1.Add("aaa", "111")
	sort.Sort(p1)
	p2 := ([]Param)(*p1)
	if !reflect.DeepEqual(p2, []Param{Param{"aaa", "111"}, Param{"bbb", "222"}}) {
		t.Fatalf("Test_SortParams err %v\n", p1)
	}
}

func TestParams_Rename(t *testing.T) {
	p := NewParams()
	p.Add("bbb", "44444")
	p.Rename("bbb", "ccc")

	if !reflect.DeepEqual(p, NewParams().Add("ccc", "44444")) {
		t.Fatalf("Test_SortParams err %v\n", p)
	}
}
