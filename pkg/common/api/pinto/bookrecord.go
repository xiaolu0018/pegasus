package pinto

import (
	"192.168.199.199/bjdaos/pegasus/pkg/common/types"
	"database/sql"
	"fmt"
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
	sqlStr := fmt.Sprint("INSERT INTO book_record(bookno,examination_no,truename,sex,bookid,bookidtype,booktimestamp,birthday,bookorg_code,createtime,telphone,book_code)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)")
	if _, err := db.Exec(sqlStr, b.BookNo, b.ExaminationNo, b.Truename, b.Sex, b.Bookid, b.Bookidtype, b.Booktimestamp, b.BirthDay, b.BookorgCode, b.CreateTime, b.Telphone, b.BookCode); err != nil {
		return err
	}
	return nil
}
