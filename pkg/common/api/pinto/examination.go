package pinto

import (
	"fmt"
	"time"

	"strconv"

	"database/sql"
	"github.com/golang/glog"

	"192.168.199.199/bjdaos/pegasus/pkg/common/types"
	"192.168.199.199/bjdaos/pegasus/pkg/pinto/server/db"
)

func GetExam(db *sql.DB, exam_no string) (*types.Examination, error) {
	sqlStr := fmt.Sprintf("SELECT status,person_code FROM examination WHERE examination_no = '%s'", exam_no)
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
	sql_number := 0
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
	sqlStr := fmt.Sprint("INSERT INTO examination_checkup(examination_no,checkup_code,checkup_status,createtime,hos_code) VALUES($1,$2,$3,$4,$5)")
	_, err := db.Exec(sqlStr, exam_checkup.ExaminationNo, exam_checkup.CheckupCode, exam_checkup.CheckupStatus, exam_checkup.CreateTime, exam_checkup.HosCode)
	return err
}

func InsertExam_Sale(db *sql.DB, exam_sale types.ExaminationSale) error {
	sqlStr := fmt.Sprint("INSERT INTO examination_checkup(examination_no,sale_code,sale_status,hos_code,sale_sellprice,discount,curprice) VALUES($1,$2,$3,$4,$5,$6,$7)")
	_, err := db.Exec(sqlStr, exam_sale.ExaminationNo, exam_sale.SaleCode, exam_sale.SaleStatus, exam_sale.HosCode, exam_sale.SaleSellprice, exam_sale.Discount, exam_sale.Curprice)
	return err
}

func MapToExams(result map[string]interface{}) (exam types.Examination, exam_checkups []types.ExaminationCheckUp, exam_sales []types.ExaminationSale, br types.BookRecord, person types.Person) {

	br = MapToBookRecord(result)
	operatetime, ok := result["operatetime"]

	operTime := time.Unix(int64(operatetime.(float64)), 0)
	if ok {
		br.CreateTime = operTime.Format("2006-01-02")
	}
	br.BookNo = operTime.Format("20060102150405")

	//person
	person.HosCode = br.BookorgCode
	person.Name = br.Truename
	person.CellPhone = br.Telphone
	person.IdcardTypeCode = br.Bookidtype
	person.Sex = br.Sex
	if marry, ok := result["merrystatus"]; ok {
		person.IsMarry = MarryToCode[marry.(string)]
	}
	person.CardNo = br.Bookid
	person.CreateTime = operTime.Format("2006-01-02 15:04:05")
	person.PersonCode = operTime.Format("20060102150405")

	//examination
	exam.CreateTime = operTime.Format("2006-01-02 15:04:05")
	exam.HosCode = br.BookorgCode
	exam.Status = "1005"
	exam.GuidePaperState = "0"
	exam.CheckupHoscode = br.BookorgCode
	exam.CheckupDate = br.Booktimestamp
	exam.ExaminationNo = GetExaminationNo(db.GetReadDB(), exam)
	exam.PersonCode = person.PersonCode

	br.ExaminationNo = exam.ExaminationNo

	//examinationCheckup
	var exam_checkup types.ExaminationCheckUp
	var exam_sale types.ExaminationSale
	if sale_codes, ok := result["sale_codes"]; ok {
		fmt.Println("sale_code ok ", sale_codes)

		sale_codesstrings := make([]string, 0, len(sale_codes.([]interface{})))

		for _, sale := range sale_codes.([]interface{}) {
			sale_codesstrings = append(sale_codesstrings, sale.(string))
		}

		if checkups, err := GetCheckupCodesBySaleCodes(db.GetReadDB(), sale_codesstrings); err == nil {
			exam_checkup.CreateTime = br.CreateTime
			exam_checkup.HosCode = br.BookorgCode
			for _, checkup := range checkups {
				exam_checkup.CheckupCode = checkup
			}
			exam_checkup.ExaminationNo = exam.ExaminationNo
			exam_checkup.CheckupStatus = 0

			exam_checkups = append(exam_checkups, exam_checkup)
		} else {
			glog.Warning("pinto.MapToExams checkups err ", err)
		}

		if sales, err := GetSalesBySaleCodes(db.GetReadDB(), sale_codesstrings); err == nil {
			exam_sale.ExaminationNo = exam.ExaminationNo
			exam_sale.HosCode = br.BookorgCode
			exam_sale.SaleStatus = "1020"
			for _, sale := range sales {
				exam_sale.SaleCode = sale.Sale_Code
				exam_sale.Discount = sale.Sale_Discount
				exam_sale.SaleSellprice = sale.Sale_SellPrice
				exam_sale.Curprice = exam_sale.Discount * exam_sale.SaleSellprice / 100
			}
			exam_sales = append(exam_sales, exam_sale)
		} else {
			glog.Warning("pinto.MapToExams exam_sale err ", err)
		}

	}

	return
}

func SaveExaminations(db *sql.DB, exam types.Examination, exam_checkups []types.ExaminationCheckUp, exam_sales []types.ExaminationSale, b types.BookRecord, person types.Person) (err error) {
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

	sqlStr := fmt.Sprint("INSERT INTO person(sex,card_no,is_marry,name,cellphone,createtime,person_code,idcard_type_code,hos_code)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)")
	_, err = tx.Exec(sqlStr, person.Sex, person.CardNo, person.IsMarry, person.Name, person.CellPhone, person.CreateTime, person.PersonCode, person.IdcardTypeCode, person.HosCode)
	if err != nil {
		glog.Error("pinto.SaveExaminations person err ", err)
		return
	}

	sqlStr = fmt.Sprint("INSERT INTO examination (examination_no,createtime,updatetime,status,person_code,org_code,hos_code,checkupdate,checkup_hoscode,guide_paper_state,report_grant_type)" +
		" VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) ON CONFLICT(examination_no) DO UPDATE SET updatetime=EXCLUDED.updatetime,status=EXCLUDED.status,person_code=EXCLUDED.person_code,org_code=EXCLUDED.org_code," +
		"hos_code=EXCLUDED.hos_code,checkupdate=EXCLUDED.checkupdate,checkup_hoscode=EXCLUDED.checkup_hoscode,guide_paper_state=EXCLUDED.guide_paper_state,report_grant_type=EXCLUDED.report_grant_type; ")
	if _, err = tx.Exec(sqlStr, exam.ExaminationNo, exam.CreateTime, exam.UpdateTime, exam.Status, exam.PersonCode, exam.OrgCode, exam.HosCode, exam.CheckupDate, exam.CheckupHoscode, exam.GuidePaperState, exam.ReportGrantType); err != nil {
		glog.Error("pinto.SaveExaminations examination err ", err)
		return
	}

	for _, exam_checkup := range exam_checkups {
		sqlStr = fmt.Sprint("INSERT INTO examination_checkup(examination_no,checkup_code,checkup_status,createtime,hos_code) VALUES($1,$2,$3,$4,$5)")
		if _, err = tx.Exec(sqlStr, exam_checkup.ExaminationNo, exam_checkup.CheckupCode, exam_checkup.CheckupStatus, exam_checkup.CreateTime, exam_checkup.HosCode); err != nil {
			glog.Error("pinto.SaveExaminations examination_checkup err ", err)
			return
		}
	}

	for _, exam_sale := range exam_sales {
		sqlStr = fmt.Sprint("INSERT INTO examination_sale(examination_no,sale_code,sale_status,hos_code,sale_sellprice,discount,curprice) VALUES($1,$2,$3,$4,$5,$6,$7)")
		if _, err = tx.Exec(sqlStr, exam_sale.ExaminationNo, exam_sale.SaleCode, exam_sale.SaleStatus, exam_sale.HosCode, exam_sale.SaleSellprice, exam_sale.Discount, exam_sale.Curprice); err != nil {
			glog.Error("pinto.SaveExaminations examination_sale err ", err)
			return
		}
	}

	sqlStr = fmt.Sprint("INSERT INTO book_record(bookno,examination_no,truename,sex,bookid,bookidtype,booktimestamp,birthday,bookorg_code,createtime,telphone,book_code)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)")
	if _, err = tx.Exec(sqlStr, b.BookNo, b.ExaminationNo, b.Truename, b.Sex, b.Bookid, b.Bookidtype, b.Booktimestamp, b.BirthDay, b.BookorgCode, b.CreateTime, b.Telphone, b.BookCode); err != nil {
		glog.Error("pinto.SaveExaminations book_record err ", err)
		return
	}

	//同时更新serial_number
	sqlStr = fmt.Sprintf("UPDATE serial_number SET seq_number = seq_number+1 WHERE hos_code = '%s' AND code = '001'", b.BookorgCode)
	if _, err = tx.Exec(sqlStr); err != nil {
		return
	}

	return
}
