package vote

import "testing"

func TestDB_Init(t *testing.T) {
	var err error
	dbI, err = NewDBInterface("192.168.199.216", "5432", "postgres", "postgresql2016", "pinto")
	if err != nil {
		t.Fatal(err)
	}

	if err = dbI.Init(); err != nil {
		t.Fatal(err)
	}
}

func TestDB_Register(t *testing.T) {
	var err error
	dbI, err = NewDBInterface("192.168.199.216", "5432", "postgres", "postgresql2016", "pinto")
	if err != nil {
		t.Fatal(err)
	}

	openid := "888888888"
	v := Voter{
		Name:        "张三",
		Image:       "image1",
		Company:     "北京迪安",
		Mobile:      "177",
		Declaration: "I can",
	}

	if err := dbI.Register(openid, &v); err != nil {
		t.Fatal(err)
	}
}

func TestDB_Register2(t *testing.T) {
	var err error
	dbI, err = NewDBInterface("192.168.199.216", "5432", "postgres", "postgresql2016", "pinto")
	if err != nil {
		t.Fatal(err)
	}

	openid := "888888889"
	v := Voter{
		Name:        "李四",
		Image:       "image1",
		Company:     "北京迪安",
		Mobile:      "189",
		Declaration: "I can",
	}

	if err := dbI.Register(openid, &v); err != nil {
		t.Fatal(err)
	}
}

func TestDB_Vote(t *testing.T) {
	var err error
	dbI, err = NewDBInterface("192.168.199.216", "5432", "postgres", "postgresql2016", "pinto")
	if err != nil {
		t.Fatal(err)
	}

	openid := "888888888"
	voterid := "888888889"
	if err := dbI.Vote(openid, voterid); err != nil {
		t.Fatal(err)
	}
}

func TestDB_ListVoters(t *testing.T) {
	var err error
	dbI, err = NewDBInterface("192.168.199.216", "5432", "postgres", "postgresql2016", "pinto")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := dbI.ListVoters(1, 10); err != nil {
		t.Fatal(err)
	}
}
