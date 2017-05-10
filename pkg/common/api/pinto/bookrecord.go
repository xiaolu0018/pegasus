package pinto

import (
	"fmt"
	"strconv"

	"database/sql"

	"bjdaos/pegasus/pkg/common/types"
	"bjdaos/pegasus/pkg/common/util/timeutil"
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
	sqlStr := fmt.Sprint(`INSERT INTO book_record(bookno,examination_no,truename,sex,bookid,bookidtype,booktimestamp,birthday,bookorg_code,createtime,telphone,book_code,is_valid)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,1)ON CONFLICT(bookno)DO UPDATE SET booktimestamp=EXCLUDED.booktimestamp`)
	_, err := db.Exec(sqlStr, b.BookNo, b.ExaminationNo, b.Truename, b.Sex, b.Bookid, b.Bookidtype, b.Booktimestamp, b.BirthDay, b.BookorgCode, b.CreateTime, b.Telphone, b.BookCode)
	return err

}

func UpdateBookRecordToInvalid(db *sql.DB, bookno string) error {
	sqlStr := fmt.Sprintf("UPDATE book_record SET is_valid = 0 WHERE bookno = '%s'", bookno)
	_, err := db.Exec(sqlStr)
	return err
}

func UpdateBookRecordWithExamToInvalid(db *sql.DB, bookno string) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	sqlStr := fmt.Sprintf("UPDATE book_record SET is_valid = 0 WHERE bookno = '%s'", bookno)
	if _, err = tx.Exec(sqlStr); err != nil {
		return
	}

	sqlStr = fmt.Sprintf("UPDATE examination SET status = 1999 WHERE examination.examination_no = (SELECT examination_no FROM book_record WHERE bookno = '%s')", bookno)
	if _, err = tx.Exec(sqlStr); err != nil {
		return
	}
	return
}

func FilterBookRecordByAppoint(a *Appointment) *types.BookRecord {
	var b types.BookRecord
	b.AppointChannel = a.Appoint_Channel
	b.BookCode = strconv.Itoa(a.AppointedNum)
	b.Bookid = a.CardNo
	b.Bookidtype = IdCardToCode[a.CardType]
	b.Sex = SexToCode[a.Sex]
	b.CreateTime = a.TimeNow.Format(timeutil.FROMAT_DAY)
	b.BookNo = a.TimeNow.Format(timeutil.FROMAT_YYMMDDHHMMSS)
	b.Truename = a.Appointor
	b.Booktimestamp = a.AppointDate
	b.Telphone = a.Mobile
	return &b
}
