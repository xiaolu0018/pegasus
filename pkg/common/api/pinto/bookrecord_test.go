package pinto

import (
	"192.168.199.199/bjdaos/pegasus/pkg/common/types"
	"192.168.199.199/bjdaos/pegasus/pkg/common/util/database"
	"fmt"
	"testing"
)

func TestGetBookRecordByBookNO(t *testing.T) {
	db, err := database.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto")
	fmt.Println("err ______", err)
	bookrecord, err := GetBookRecordByBookNO(db, "2017042913523149951")
	fmt.Println("err __", err, bookrecord)
}

func TestInsertBookRecord(t *testing.T) {
	fmt.Println("enter")
	db, err := database.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto")
	fmt.Println("err__", err, db)
	b := types.BookRecord{
		BookCode:       "3",
		BookorgCode:    "0001001",
		BookNo:         "201704291818",
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

	err = InsertBookRecord(db, &b)
	fmt.Println("err", err)
}
