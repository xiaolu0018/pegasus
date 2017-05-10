package pinto

import (
	"bjdaos/pegasus/pkg/common/types"
	"bjdaos/pegasus/pkg/common/util/database"
	"testing"
)

func TestGetBookRecordByBookNO(t *testing.T) {
	db, _ := database.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto")
	if _, err := GetBookRecordByBookNO(db, "2017042913523149951"); err != nil {
		t.Fatal(err)
	}
}

func TestInsertBookRecord(t *testing.T) {
	db, err := database.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto")
	b := types.BookRecord{
		BookCode:       "3",
		BookorgCode:    "0001001",
		BookNo:         "20170509193904",
		Truename:       "kang",
		Telphone:       "19865441259",
		Sex:            1,
		Bookid:         "610488888888888888",
		Booktimestamp:  "2017-05-05",
		BirthDay:       "",
		AppointChannel: "微信",
		CreateTime:     "2017-04-29 18:22:55",
		Bookidtype:     "3",
	}

	if err = InsertBookRecord(db, &b); err != nil {
		t.Fatal(err)
	}

}
