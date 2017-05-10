package vote

import "testing"

func TestDB_Init(t *testing.T) {
	var err error
	DBI, err = NewDBInterface("10.1.0.190", "5432", "postgres", "postgres190@", "pinto")
	if err != nil {
		t.Fatal(err)
	}
	var access_token string = "SO0okab3DtJF1I1lkWwN5YUP1wn9YyVqVZ98xVMD2LQkc_P5sUk-RdpVtnbsNWN1pw9smf94Dauqoaa1cVvLbyNUzIM5VKK7NiK53UlYEk5d6jY_iw58Q8ODifFox_eODJLbABAXDF"
	if err = DBI.Init(access_token); err != nil {
		t.Fatal(err)
	}
}

func TestDB_Register(t *testing.T) {
	var err error
	DBI, err = NewDBInterface("10.1.0.190", "5432", "postgres", "postgres190@", "pinto")
	if err != nil {
		t.Fatal(err)
	}

	v := Voter{
		OpenID: "888888888",
		Name:        "张三",
		Image:       "image1",
		Company:     "北京迪安",
		Mobile:      "177",
		Declaration: "I can",
	}

	if err := DBI.Register( &v); err != nil {
		t.Fatal(err)
	}
}

func TestDB_Register2(t *testing.T) {
	var err error
	DBI, err = NewDBInterface("192.168.199.216", "5432", "postgres", "postgresql2016", "pinto")
	if err != nil {
		t.Fatal(err)
	}

	v := Voter{
		OpenID : "888888889",
		Name:        "李四",
		Image:       "image1",
		Company:     "北京迪安",
		Mobile:      "189",
		Declaration: "I can",
	}

	if err := DBI.Register(&v); err != nil {
		t.Fatal(err)
	}
}

func TestDB_Vote(t *testing.T) {
	var err error
	DBI, err = NewDBInterface("192.168.199.216", "5432", "postgres", "postgresql2016", "pinto")
	if err != nil {
		t.Fatal(err)
	}

	openid := "888888888"
	voterid := "888888889"
	if err := DBI.Vote(openid, voterid); err != nil {
		t.Fatal(err)
	}
}

func TestDB_ListVoters(t *testing.T) {
	var err error
	DBI, err = NewDBInterface("192.168.199.216", "5432", "postgres", "postgresql2016", "pinto")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := DBI.ListVoters(nil,1, 10); err != nil {
		t.Fatal(err)
	}
}
