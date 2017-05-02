package pinto

import (
	"192.168.199.199/bjdaos/pegasus/pkg/common/types"
	"database/sql"
	"fmt"
	"strconv"
)

func GetExam(db *sql.DB, exam_no string) (*types.Examination, error) {
	sqlStr := fmt.Sprintf("SELECT status,Person_code FROM examination WHERE examination_no = '%s'", exam_no)
	exam := types.Examination{}
	if err := db.QueryRow(sqlStr).Scan(&exam.Status, &exam.PersonCode); err != nil {
		return nil, err
	}
	return &exam, nil
}

func CreateUpdate(db *sql.DB, exam types.Examination) error {
	sqlStr := fmt.Sprint("INSERT INTO examination (examination_no,createtime,updatetime,status,person_code,org_code,hos_code,checkupdate,checkup_hoscode,guide_paper_state,report_grant_type)" +
		" VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) ON CONFLICT(examination_no) DO UPDATE SET updatetime=EXCLUDED.updatetime,status=EXCLUDED.status,person_code=EXCLUDED.person_code,org_code=EXCLUDED.org_code," +
		"hos_code=EXCLUDED.hos_code,checkupdate=EXCLUDED.checkupdate,checkup_hoscode=EXCLUDED.checkup_hoscode,guide_paper_state=EXCLUDED.guide_paper_state,report_grant_type=EXCLUDED.report_grant_type; ")
	_, err := db.Exec(sqlStr, exam.ExaminationNo, exam.CreateTime, exam.UpdateTime, exam.Status, exam.PersonCode, exam.OrgCode, exam.HosCode, exam.CheckupDate, exam.CheckupHoscode, exam.GuidePaperState, exam.ReportGrantType)
	return err

}

func GetExaminationNo(db *sql.DB, exam types.Examination) string {
	str := exam.HosCode
	str += exam.CreateTime[2:4]
	sqlStr := fmt.Sprintf("SELECT seq_number FROM serial_number WHERE hos_code='%s' AND code = '%s'", exam.HosCode, "001")
	var sql_number int
	if err := db.QueryRow(sqlStr).Scan(&sql_number); err != nil {
		fmt.Println("get exam no ", err)
		return ""
	}
	sql_numberstr := strconv.Itoa(sql_number)
	var numberinit = "0000000"
	if len(sql_numberstr) < 7 {
		sql_numberstr = numberinit[0:7-len(sql_numberstr)] + sql_numberstr
	}
	str += sql_numberstr
	return str
}

func InsertExam_Checkup(db *sql.DB, exam_checkup types.ExaminationCheckUp) error {
	sqlStr := fmt.Sprintf("INSERT INTO examination_checkup(examination_no,checkup_code,checkup_status,createtime,hos_code) VALUES($1,$2,$3,$4,$5)")
	_, err := db.Exec(sqlStr, exam_checkup.ExaminationNo, exam_checkup.CheckupCode, exam_checkup.CheckupStatus, exam_checkup.CreateTime, exam_checkup.HosCode)
	return err
}

func InsertExam_Sale(db *sql.DB, exam_sale types.ExaminationSale) error {
	sqlStr := fmt.Sprintf("INSERT INTO examination_checkup(examination_no,sale_code,sale_status,hos_code,sale_sellprice,discount,curprice) VALUES($1,$2,$3,$4,$5,$6,$7)")
	_, err := db.Exec(sqlStr, exam_sale.ExaminationNo, exam_sale.SaleCode, exam_sale.SaleStatus, exam_sale.HosCode, exam_sale.SaleSellprice, exam_sale.Discount, exam_sale.Curprice)
	return err
}

func SaveExaminations(db *sql.DB, exam types.Examination, exam_checkup types.ExaminationCheckUp, exam_sale types.ExaminationSale, b types.BookRecord, person types.Person) (err error) {
	var tx *sql.Tx
	tx, err = db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	sqlStr := fmt.Sprintf("INSERT INTO person(sex,card_no,is_marry,name,cellphone,createtime,person_code,idcard_type_code,hos_code)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)")
	_, err = tx.Exec(sqlStr, person.Sex, person.CardNo, person.IsMarry, person.Name, person.CellPhone, person.CreateTime, person.PersonCode, person.IdcardTypeCode, person.HosCode)

	sqlStr = fmt.Sprint("INSERT INTO examination (examination_no,createtime,updatetime,status,person_code,org_code,hos_code,checkupdate,checkup_hoscode,guide_paper_state,report_grant_type)" +
		" VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) ON CONFLICT(examination_no) DO UPDATE SET updatetime=EXCLUDED.updatetime,status=EXCLUDED.status,person_code=EXCLUDED.person_code,org_code=EXCLUDED.org_code," +
		"hos_code=EXCLUDED.hos_code,checkupdate=EXCLUDED.checkupdate,checkup_hoscode=EXCLUDED.checkup_hoscode,guide_paper_state=EXCLUDED.guide_paper_state,report_grant_type=EXCLUDED.report_grant_type; ")
	if _, err = tx.Exec(sqlStr, exam.ExaminationNo, exam.CreateTime, exam.UpdateTime, exam.Status, exam.PersonCode, exam.OrgCode, exam.HosCode, exam.CheckupDate, exam.CheckupHoscode, exam.GuidePaperState, exam.ReportGrantType); err != nil {
		return
	}

	sqlStr = fmt.Sprintf("INSERT INTO examination_checkup(examination_no,checkup_code,checkup_status,createtime,hos_code) VALUES($1,$2,$3,$4,$5)")
	if _, err = tx.Exec(sqlStr, exam_checkup.ExaminationNo, exam_checkup.CheckupCode, exam_checkup.CheckupStatus, exam_checkup.CreateTime, exam_checkup.HosCode); err != nil {
		return
	}

	sqlStr = fmt.Sprintf("INSERT INTO examination_checkup(examination_no,sale_code,sale_status,hos_code,sale_sellprice,discount,curprice) VALUES($1,$2,$3,$4,$5,$6,$7)")
	if _, err = tx.Exec(sqlStr, exam_sale.ExaminationNo, exam_sale.SaleCode, exam_sale.SaleStatus, exam_sale.HosCode, exam_sale.SaleSellprice, exam_sale.Discount, exam_sale.Curprice); err != nil {
		return
	}

	sqlStr = fmt.Sprint("INSERT INTO book_record(bookno,examination_no,truename,sex,bookid,bookidtype,booktimestamp,birthday,bookorg_code,createtime,telphone,book_code)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)")
	if _, err = tx.Exec(sqlStr, b.BookNo, b.ExaminationNo, b.Truename, b.Sex, b.Bookid, b.Bookidtype, b.Booktimestamp, b.BirthDay, b.BookorgCode, b.CreateTime, b.Telphone, b.BookCode); err != nil {
		return
	}

	return
}
