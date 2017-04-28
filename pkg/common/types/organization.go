package types

import (
	"sort"
)

type Organization struct {
	ID         string //id
	Code       string //org_code
	Name       string //org_name
	Phone      string //org_phone
	ImageUrl   string
	DetailsUrl string
	Deleted    bool
}

type OrganizationList []Organization

func (l *OrganizationList) Len() int {
	return len(([]Organization)(*l))
}

func (l *OrganizationList) Less(i, j int) bool {
	return (*l)[i].Code <= (*l)[j].Code
}

func (l *OrganizationList) Swap(i, j int) {
	tmp := (*l)[i]
	(*l)[i] = (*l)[j]
	(*l)[j] = tmp
}

func Sort(l []Organization) []Organization {
	newL := OrganizationList(l)
	sort.Sort(&newL)
	return ([]Organization)(newL)
}

func CodeEquals(a, b Organization) bool {
	return a.Code == b.Code
}

func Equals(a, b Organization) bool {
	return a == b
}

func (a Organization) Large(b Organization) bool {
	return a.Code > b.Code
}

func Diff(new, old []Organization) (add, del, change []Organization) {
	new, old = Sort(new), Sort(old)
	i, j := 0, 0
	for {
		if i == len(new) || j == len(old) {
			if i == len(new) && j < len(old) {
				del = append(del, old[j:]...)
				return
			}
			if j == len(old) && i < len(new) {
				add = append(add, new[i:]...)
				return
			}
			if i == len(new) && j == len(old) {
				return
			}
		}

		switch {
		case CodeEquals(new[i], old[j]):
			if !Equals(new[i], old[j]) {
				change = append(change, new[i])
			}
			i++
			j++
		case new[i].Large(old[j]):
			del = append(del, old[j])
			j++
		case old[j].Large(new[i]):
			add = append(add, new[i])
			i++
		}
	}
	return
}
