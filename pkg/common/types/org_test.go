package types

import (
	"sort"
	"testing"
)

func TestOrganizationList(t *testing.T) {
	l := &OrganizationList{
		{"id0004", "code001", "name001"},
		{"id0001", "code001", "name001"},
		{"id0002", "code001", "name001"},
		{"id0003", "code001", "name001"},
	}

	sort.Sort(l)
}

func TestSort(t *testing.T) {
	l := []Organization{
		{"id0004", "code004", "name001"},
		{"id0001", "code001", "name001"},
		{"id0002", "code002", "name001"},
		{"id0003", "code003", "name001"},
	}

	newL := Sort(l)
	o := Organization{"id0001", "code001", "name001"}
	if newL[0] != o {
		t.Fatal()
	}
}

func TestDiff(t *testing.T) {
	newL := []Organization{
		{"id0001", "code001", "name001"},
	}
	oldL := []Organization{
		{"id0001", "code001", "name001"},
	}
	a, b, c := Diff(newL, oldL)
	if len(a) != 0 || len(b) != 0 || len(c) != 0 {
		t.Fatal(a, b, c)
	}

	newL = []Organization{}
	oldL = []Organization{}
	a, b, c = Diff(newL, oldL)
	if len(a) != 0 || len(b) != 0 || len(c) != 0 {
		t.Fatal(a, b, c)
	}

	newL = []Organization{
		{"id0001", "code001", "name001"},
	}
	oldL = []Organization{}
	a, b, c = Diff(newL, oldL)
	if len(a) != 1 || len(b) != 0 || len(c) != 0 {
		t.Fatal(a, b, c)
	}

	newL = []Organization{}
	oldL = []Organization{
		{"id0001", "code001", "name001"},
	}
	a, b, c = Diff(newL, oldL)
	if len(a) != 0 || len(b) != 1 || len(c) != 0 {
		t.Fatal(a, b, c)
	}

	newL = []Organization{
		{"id0001", "code001", "name001"},
		{"id0003", "code003", "name003"},
		{"id0005", "code005", "name005"},
	}
	oldL = []Organization{
		{"id0002", "code002", "name002"},
		{"id0004", "code004", "name004"},
		{"id0006", "code006", "name006"},
	}
	a, b, c = Diff(newL, oldL)
	if len(a) != 3 || len(b) != 3 || len(c) != 0 {
		t.Fatal(a, b, c)
	}

}
