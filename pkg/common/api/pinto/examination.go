package pinto

import (
	"fmt"

	"strconv"

	"database/sql"
	"github.com/golang/glog"

	"bjdaos/pegasus/pkg/common/types"
	"bjdaos/pegasus/pkg/common/util/timeutil"
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

	glog.Errorln("sql_number", sql_number)
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

func FilterExamCheckups(db *sql.DB, a *Appointment, exam_no string) ([]types.ExaminationCheckUp, error) {
	var exam_checkups []types.ExaminationCheckUp
	var exam_checkup types.ExaminationCheckUp

	if checkups, err := GetCheckupCodesBySaleCodes(db, a.SaleCodes); err == nil {
		exam_checkup.CreateTime = a.TimeNow.Format(timeutil.FROMAT_DAY)
		exam_checkup.HosCode = a.OrgCode
		for _, checkup := range checkups {
			exam_checkup.CheckupCode = checkup
		}
		exam_checkup.ExaminationNo = exam_no
		exam_checkup.CheckupStatus = 0
		exam_checkups = append(exam_checkups, exam_checkup)
	} else {
		glog.Warning("pinto.MapToExams checkups err ", err)
		return nil, err
	}
	return exam_checkups, nil
}

func FilterExamSales(db *sql.DB, a *Appointment, exam_no string) ([]types.ExaminationSale, error) {

	var exam_sales []types.ExaminationSale
	var exam_sale types.ExaminationSale

	if sales, err := GetSalesBySaleCodes(db, a.SaleCodes); err == nil {
		exam_sale.ExaminationNo = exam_no
		exam_sale.HosCode = a.OrgCode
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
		return nil, err
	}
	return exam_sales, nil
}

func FilterExam(db *sql.DB, a *Appointment, p_code string) (*types.Examination, error) {
	var exam types.Examination
	exam.HosCode = a.OrgCode
	exam.PersonCode = p_code
	exam.CreateTime = a.TimeNow.Format(timeutil.FROMAT_DAY)
	exam.CheckupDate = a.AppointDate
	exam.GuidePaperState = "0"
	exam.Status = "1020"
	exam.ReportGrantType = "0"
	exam.ExaminationNo = GetExaminationNo(db, exam)
	return &exam, nil
}

func FilterExamsAll(db *sql.DB, a *Appointment) (*ExamsAll, error) {
	var examAll ExamsAll
	examAll.B = FilterBookRecordByAppoint(a)
	examAll.P = FilterPersonByAppoint(a)
	var err error
	if examAll.E, err = FilterExam(db, a, examAll.P.PersonCode); err != nil {
		return nil, err
	}

	if examAll.Checkups, err = FilterExamCheckups(db, a, examAll.E.ExaminationNo); err != nil {
		return nil, err
	}

	if examAll.Sales, err = FilterExamSales(db, a, examAll.E.ExaminationNo); err != nil {
		return nil, err
	}
	return &examAll, err

}

func SaveExaminations(db *sql.DB, e *ExamsAll) (err error) {
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

	person := *e.P
	sqlStr := fmt.Sprint("INSERT INTO person(sex,card_no,is_marry,name,cellphone,createtime,person_code,idcard_type_code,hos_code)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)")
	_, err = tx.Exec(sqlStr, person.Sex, person.CardNo, person.IsMarry, person.Name, person.CellPhone, person.CreateTime, person.PersonCode, person.IdcardTypeCode, person.HosCode)
	if err != nil {
		glog.Error("pinto.SaveExaminations person err ", err)
		return
	}

	exam := *e.E
	sqlStr = fmt.Sprint("INSERT INTO examination (examination_no,createtime,updatetime,status,person_code,org_code,hos_code,checkupdate,checkup_hoscode,guide_paper_state,report_grant_type)" +
		" VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) ON CONFLICT(examination_no) DO UPDATE SET updatetime=EXCLUDED.updatetime,status=EXCLUDED.status,person_code=EXCLUDED.person_code,org_code=EXCLUDED.org_code," +
		"hos_code=EXCLUDED.hos_code,checkupdate=EXCLUDED.checkupdate,checkup_hoscode=EXCLUDED.checkup_hoscode,guide_paper_state=EXCLUDED.guide_paper_state,report_grant_type=EXCLUDED.report_grant_type; ")
	if _, err = tx.Exec(sqlStr, exam.ExaminationNo, exam.CreateTime, exam.UpdateTime, exam.Status, exam.PersonCode, exam.OrgCode, exam.HosCode, exam.CheckupDate, exam.CheckupHoscode, exam.GuidePaperState, exam.ReportGrantType); err != nil {
		glog.Error("pinto.SaveExaminations examination err ", err)
		return
	}

	exam_checkups := e.Checkups
	for _, exam_checkup := range exam_checkups {
		sqlStr = fmt.Sprint("INSERT INTO examination_checkup(examination_no,checkup_code,checkup_status,createtime,hos_code) VALUES($1,$2,$3,$4,$5)")
		if _, err = tx.Exec(sqlStr, exam_checkup.ExaminationNo, exam_checkup.CheckupCode, exam_checkup.CheckupStatus, exam_checkup.CreateTime, exam_checkup.HosCode); err != nil {
			glog.Error("pinto.SaveExaminations examination_checkup err ", err)
			return
		}
	}

	exam_sales := e.Sales
	for _, exam_sale := range exam_sales {
		sqlStr = fmt.Sprint("INSERT INTO examination_sale(examination_no,sale_code,sale_status,hos_code,sale_sellprice,discount,curprice) VALUES($1,$2,$3,$4,$5,$6,$7)")
		if _, err = tx.Exec(sqlStr, exam_sale.ExaminationNo, exam_sale.SaleCode, exam_sale.SaleStatus, exam_sale.HosCode, exam_sale.SaleSellprice, exam_sale.Discount, exam_sale.Curprice); err != nil {
			glog.Error("pinto.SaveExaminations examination_sale err ", err)
			return
		}
	}

	b := *e.B
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
