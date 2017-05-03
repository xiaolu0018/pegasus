package pinto

import (
	"fmt"
	"time"

	"database/sql"

	"192.168.199.199/bjdaos/pegasus/pkg/common/types"
	"strconv"
)

func GetBookRecordByBookNO(db *sql.DB, bookno string) (*types.BookRecord, error) {
	sqlStr := fmt.Sprintf("SELECT examination_no FROM book_record WHERE bookno = '%s'", bookno)
	var bookrecord types.BookRecord
	if err := db.QueryRow(sqlStr).Scan(&bookrecord.ExaminationNo); err != nil {
		return nil, err
	}
	bookrecord.BookNo = bookno
	return &bookrecord, nil
}

func InsertBookRecord(db *sql.DB, b *types.BookRecord) error {
	sqlStr := fmt.Sprint("INSERT INTO book_record(bookno,examination_no,truename,sex,bookid,bookidtype,booktimestamp,birthday,bookorg_code,createtime,telphone,book_code,is_valid)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,1)")
	if _, err := db.Exec(sqlStr, b.BookNo, b.ExaminationNo, b.Truename, b.Sex, b.Bookid, b.Bookidtype, b.Booktimestamp, b.BirthDay, b.BookorgCode, b.CreateTime, b.Telphone, b.BookCode); err != nil {
		return err
	}
	return nil
}

func MapToBookRecord(result map[string]interface{}) types.BookRecord {
	br := types.BookRecord{}
	if cartno, ok := result["cardno"]; ok {
		br.Bookid = cartno.(string)
	}
	if cardtype, ok := result["cardtype"]; ok {
		br.Bookidtype = IdCardToCode[cardtype.(string)]
	}

	if appointor, ok := result["appointor"]; ok {
		br.Truename = appointor.(string)
	}
	if sex, ok := result["sex"]; ok {
		br.Sex = SexToCode[sex.(string)]
	}

	if appointtime, ok := result["appointtime"]; ok {
		br.Booktimestamp = time.Unix(int64(appointtime.(float64)), 0).Format("2006-01-02")

	}

	operatetime, ok := result["operatetime"]

	operTime := time.Unix(int64(operatetime.(float64)), 0)
	if ok {
		br.CreateTime = operTime.Format("2006-01-02")
	}

	if org_code, ok := result["org_code"]; ok {
		br.BookorgCode = org_code.(string)
	}

	if mobile, ok := result["mobile"]; ok {
		br.Telphone = mobile.(string)
	}

	if AppointedNum, ok := result["appointednum"]; ok {
		br.BookCode = strconv.FormatInt(int64(AppointedNum.(float64)), 10)
	}

	br.BookNo = operTime.Format("20060102150405")
	return br
}
