package vote

import (
	"testing"
	"sort"
	"fmt"
)

func TestDB_GetVoter(t *testing.T) {
	r := []int64{1,23,234235,626345}
	sort.Sort(Int64Slice(r))
	fmt.Println(r)
}

//func TestDB_GetVoter(t *testing.T) {
//	var err error
//	DBI, err = NewDBInterface("10.1.0.190", "5432", "postgres", "postgres190@", "pinto")
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	v, err := DBI.GetVoter("123")
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println(v)
//}

func TestDB_Register(t *testing.T) {
	var err error
	DBI, err = NewDBInterface("10.1.0.190", "5432", "postgres", "postgres190@", "pinto")
	if err != nil {
		t.Fatal(err)
	}

	v := &Voter{}
	v.OpenID = "123"
	v.Name = "michael"
	v.Image = "456.jpg"
	v.Company = "baidu"
	v.Mobile = "177"
	v.Declaration = "111123123"
	v.imageCached = true

	if err := DBI.Register(v); err != nil {
		t.Fatal(err)
	}
}

//
//func TestDB_Init(t *testing.T) {
//	var err error
//	DBI, err = NewDBInterface("10.1.0.190", "5432", "postgres", "postgres190@", "pinto")
//	if err != nil {
//		t.Fatal(err)
//	}
//	var access_token string = "SO0okab3DtJF1I1lkWwN5YUP1wn9YyVqVZ98xVMD2LQkc_P5sUk-RdpVtnbsNWN1pw9smf94Dauqoaa1cVvLbyNUzIM5VKK7NiK53UlYEk5d6jY_iw58Q8ODifFox_eODJLbABAXDF"
//	if err = DBI.Init(access_token); err != nil {
//		t.Fatal(err)
//	}
//}
//
//func TestDB_Register(t *testing.T) {
//	var err error
//	DBI, err = NewDBInterface("10.1.0.190", "5432", "postgres", "postgres190@", "pinto")
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	v := Voter{
//		OpenID: "888888888",
//		Name:        "张三",
//		Image:       "image1",
//		Company:     "北京迪安",
//		Mobile:      "177",
//		Declaration: "I can",
//	}
//
//	if err := DBI.Register( &v); err != nil {
//		t.Fatal(err)
//	}
//}
//
//func TestDB_Register2(t *testing.T) {
//	var err error
//	DBI, err = NewDBInterface("192.168.199.216", "5432", "postgres", "postgresql2016", "pinto")
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	v := Voter{
//		OpenID : "888888889",
//		Name:        "李四",
//		Image:       "image1",
//		Company:     "北京迪安",
//		Mobile:      "189",
//		Declaration: "I can",
//	}
//
//	if err := DBI.Register(&v); err != nil {
//		t.Fatal(err)
//	}
//}
//
//func TestDB_Vote(t *testing.T) {
//	var err error
//	DBI, err = NewDBInterface("192.168.199.216", "5432", "postgres", "postgresql2016", "pinto")
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	openid := "888888888"
//	voterid := "888888889"
//	if err := DBI.Vote(openid, voterid); err != nil {
//		t.Fatal(err)
//	}
//}
//
//func TestDB_ListVoters(t *testing.T) {
//	var err error
//	DBI, err = NewDBInterface("192.168.199.216", "5432", "postgres", "postgresql2016", "pinto")
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	if _, err := DBI.ListVoters(nil,1, 10); err != nil {
//		t.Fatal(err)
//	}
//}
