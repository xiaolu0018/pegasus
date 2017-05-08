package appointment

import (
	"fmt"
	"time"

	"strings"

	"database/sql"
	"encoding/json"
	"github.com/lib/pq"

	"github.com/golang/glog"

	"bjdaos/pegasus/pkg/appoint/db"
	org "bjdaos/pegasus/pkg/appoint/organization"
	"bjdaos/pegasus/pkg/common/util/methods"
)

/*
	0, 查Org_SpecialItem 得到所选项目的限制人数
	1，查org_config  得到 容量与报警数值
	2，查managercapacity 判断预约数，与 1 中的 容量和报警数值 比较
		正常 预约成功，更新 managercapacity，不正常报错
	3，查managerItem  判断预约数，与 0 中的限制人数 比较
		正常 预约成功，更新 managerItem，不正常报错
	4，保存 a
*/
func (a *Appointment) CreateAppointment() (err error) {
	if a.ID == "" {
		a.ID = time.Now().Format("20060102150405999")
	}

	if total := GetAppointmentByCardNo(a.CardNo, a.AppointTime); total > 0 {
		return fmt.Errorf("This user info had appointed!")
	}

	tx, err := db.GetDB().Begin()
	if err != nil {
		return err
	}

	err = addAppointment(tx, a)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}

	if a.PlanId != "" {
		return sendToPintoWithPlan(a)
	} else {
		return sendToPintoWithOutPlan(a)
	}
}

func sendToPintoWithOutPlan(a *Appointment) error {
	result, statuscode, errs := methods.Go_Through_HttpWithBody("POST", "http://192.168.199.198:9300", "/api/appointment/person", "", a)
	if errs != nil || statuscode != 200 {
		for i := 0; i < 3; i++ {
			if result, statuscode, errs = methods.Go_Through_HttpWithBody("POST", "http://192.168.199.198:9300", "/api/appointment/person", "", a); errs == nil && statuscode == 200 {
				break
			}
		}
	}
	resultmap := map[string]string{}

	json.Unmarshal(result, &resultmap)

	sqlStr := fmt.Sprintf("UPDATE %s SET bookno = '%s' WHERE id = '%s'", T_APPOINTMENT, resultmap["bookno"], a.ID)
	if _, errs := db.GetDB().Exec(sqlStr); errs != nil {
		return errs
	}
	return nil
}

func sendToPintoWithPlan(a *Appointment) error {
	result, statuscode, errs := methods.Go_Through_HttpWithBody("POST", "http://192.168.199.198:9300", "/api/appointment/plan", "", a)
	if errs != nil || statuscode != 200 {
		for i := 0; i < 3; i++ {
			if result, statuscode, errs = methods.Go_Through_HttpWithBody("POST", "http://192.168.199.198:9300", "/api/appointment/plan", "", a); errs == nil && statuscode == 200 {
				break
			}
		}
	}
	resultmap := map[string]string{}

	json.Unmarshal(result, &resultmap)

	sqlStr := fmt.Sprintf("UPDATE %s SET bookno = '%s' WHERE id = '%s'", T_APPOINTMENT, resultmap["bookno"], a.ID)
	if _, errs := db.GetDB().Exec(sqlStr); errs != nil {
		return errs
	}
	return nil
}

func (a *Appointment) CancelAppointment() (err error) {
	tx := &sql.Tx{}
	if tx, err = db.GetDB().Begin(); err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	//设置appointment  cancel = 'true'
	sql := fmt.Sprintf(`UPDATE %s SET ifcancel = 'true',status = '已取消' WHERE id = '%s'`, T_APPOINTMENT, a.ID)
	if _, err = tx.Exec(sql); err != nil {
		glog.Errorln("CancelAppointment update appointment err", err.Error())
		return
	}
	err = deleteAppointment(tx, a)

	var errs error

	if errs = sentToPintoForCancel(a); errs != nil {
		glog.Errorln("appoint.CancelAppointment sentToPintoForCancel err", errs.Error())
	}

	return
}

func sentToPintoForCancel(a *Appointment) error {
	basic, err := org.GetConfigBasic(a.OrgCode)
	if err != nil {
		return err
	}
	result := make(map[string]interface{})
	result["bookno"] = a.BookNo
	result["withplan"] = false
	if a.PlanId != "" {
		result["withplan"] = true
	}

	if _, statuscode, errs := methods.Go_Through_HttpWithBody("POST", basic.IpAddress, "/api/examination/cancel", "", result); errs == nil && statuscode == 200 {
		return nil
	} else {
		for i := 0; i < 3; i++ {
			if _, statuscode, errs = methods.Go_Through_HttpWithBody("POST", basic.IpAddress, "/api/examination/cancel", "", result); errs == nil && statuscode == 200 {
				break
			}
		}
		return errs
	}
}

func (a *Appointment) UpdateAppointment() (err error) {
	tx := &sql.Tx{}
	if tx, err = db.GetDB().Begin(); err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	oldAppointment := &Appointment{}
	if oldAppointment, err = GetAppointment(a.ID); err != nil {
		return
	}

	if err = deleteAppointment(tx, oldAppointment); err != nil {
		return
	}
	if err = addAppointment(tx, a); err != nil {
		return
	}

	return
}

func GetAppointment(appointid string) (*Appointment, error) {
	sqlStr := fmt.Sprintf("SELECT id,appointtime,org_code,planid,cardtype,cardno,mobile,appointor,merrystatus,status,appoint_channel,"+
		`company,"group",remark,operator,operatetime,orderid,commentid,appointednum,ifsingle,ifcancel,sales_code,bookno FROM %s WHERE id = '%s'`, T_APPOINTMENT, appointid)

	var id, org_code, planid, cardtype, cardno, mobile, appointor, merrystatus, status, appoint_channel, company, group, remark, operator, orderid, commentid, bookno string
	var appointtime, operatetime int64
	var appointednum int
	var ifsingle, ifcancel bool
	var salescode pq.StringArray
	err := db.GetDB().QueryRow(sqlStr).Scan(&id, &appointtime, &org_code, &planid, &cardtype, &cardno, &mobile, &appointor, &merrystatus, &status, &appoint_channel,
		&company, &group, &remark, &operator, &operatetime, &orderid, &commentid, &appointednum, &ifsingle, &ifcancel, salescode, &bookno)
	if err != nil {
		return nil, err
	}
	a := Appointment{
		ID:              id,
		OrgCode:         org_code,
		AppointTime:     appointtime,
		PlanId:          planid,
		SaleCodes:       []string(salescode),
		CardType:        cardtype,
		CardNo:          cardno,
		Mobile:          mobile,
		Appointor:       appointor,
		MerryStatus:     merrystatus,
		Status:          status,
		Appoint_Channel: appoint_channel,
		Company:         company,
		Group:           group,
		Remark:          remark,
		Operator:        operator,
		OperateTime:     operatetime,
		OrderID:         orderid,
		CommentID:       commentid,
		AppointedNum:    appointednum,
		IfSingle:        ifsingle,
		IfCancel:        ifcancel,
		BookNo:          bookno,
	}
	return &a, nil
}

func GetAppointmentByCardNo(cardno string, appointtime int64) int {
	sqlStr := fmt.Sprintf("SELECT count(*) FROM %s WHERE cardno = '%s' AND appointtime = %d", T_APPOINTMENT, cardno, appointtime)
	var count int
	if err := db.GetDB().QueryRow(sqlStr).Scan(&count); err != nil {
		glog.Errorln("appoint.GetAppointmentByCardNo err ", err)
		return 0
	}
	return count
}

func GetAppointmentList(page_index, page_size int, begintime, endtime int64, org_code, search, userid string) ([]Appointment, int, error) {
	var totalnums int
	apps := make([]Appointment, 0, page_size)
	if len(org_code) != 0 {
		org_code = fmt.Sprintf(`AND org_code = '%s'`, org_code)
	}
	var beginTimeSql, endTimeSql string
	beginTimeSql = fmt.Sprintf("appointtime >= '%d' ", begintime)

	if endtime > 0 {
		endTimeSql = fmt.Sprintf("AND appointtime <= '%d'", endtime)
	}

	if len(search) != 0 {
		search = fmt.Sprintf("AND ( appointor = '%s' OR mobile = '%s' OR cardno = '%s'", search, search, search)
	}

	if userid != "" {
		userid = fmt.Sprintf("AND appointorid = '%s'", userid)
	}
	if err := db.GetDB().QueryRow(fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE %s %s %s %s %s`,
		T_APPOINTMENT, beginTimeSql, search, org_code, endTimeSql, userid)).Scan(&totalnums); err != nil {
		glog.Errorf("GetQueryAll: sql  err %v\n", err)
		return nil, 0, err
	}
	sqlStr := fmt.Sprintf("SELECT id,appointtime,org_code,planid,cardtype,cardno,mobile,appointor,appointorid,merrystatus,status,appoint_channel,"+
		`company,"group",remark,operator,operatetime,orderid,commentid,appointednum,reportid,ifsingle,ifcancel FROM %s WHERE  %s %s %s %s %s LIMIT '%d' OFFSET '%d' `,
		T_APPOINTMENT, beginTimeSql, search, org_code, endTimeSql, userid, page_size, page_index)

	var id, orgcode, planid, cardtype, cardno, mobile, appointor, appointorid, merrystatus, status,
		appoint_channel, company, group, remark, operator, orderid, commentid, reportid string
	var appointtime, operatetime int64
	var appointednum int
	var ifsingle, ifcancel bool
	var rows *sql.Rows
	var err error
	if rows, err = db.GetDB().Query(sqlStr); err != nil {
		glog.Errorln("appointment.GetAppointmentList Query,err " + err.Error())
		return nil, 0, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&id, &appointtime, &orgcode, &planid, &cardtype, &cardno, &mobile, &appointor, &appointorid, &merrystatus, &status, &appoint_channel,
			&company, &group, &remark, &operator, &operatetime, &orderid, &commentid, &appointednum, &reportid, &ifsingle, &ifcancel); err != nil {
			glog.Errorln("appointment.GetAppointmentList  rows.Scan,err " + err.Error())
			return nil, 0, err
		}
		a := Appointment{
			ID:              id,
			OrgCode:         orgcode,
			AppointTime:     appointtime,
			PlanId:          planid,
			CardType:        cardtype,
			CardNo:          cardno,
			Mobile:          mobile,
			Appointor:       appointor,
			Appointorid:     appointorid,
			MerryStatus:     merrystatus,
			Status:          status,
			Appoint_Channel: appoint_channel,
			Company:         company,
			Group:           group,
			Remark:          remark,
			Operator:        operator,
			OperateTime:     operatetime,
			OrderID:         orderid,
			CommentID:       commentid,
			AppointedNum:    appointednum,
			ReportId:        reportid,
			IfSingle:        ifsingle,
			IfCancel:        ifcancel,
		}

		apps = append(apps, a)
	}

	if rows.Err() != nil {
		return nil, 0, rows.Err()
	}

	for k, app := range apps {
		app.AppointDate = time.Unix(app.AppointTime, 0).Format("2006-01-02")
		app.OperateDate = time.Unix(app.OperateTime, 0).Format("2006-01-02")
		apps[k] = app
	}

	return apps, totalnums, nil
}

func addAppointment(tx *sql.Tx, app *Appointment) (err error) {
	basicCfg, err := org.GetConfigBasic(app.OrgCode)
	if err != nil {
		glog.Errorf("appoint.addAppointment GetConfigBasic err ", err.Error())
		return
	}

	app.AppointDate = time.Unix(app.AppointTime, 0).Format("2006-01-02")
	capUsed, err := getAppDateRecordNum(tx, app.OrgCode, app.AppointDate)
	if err != nil {
		fmt.Println("xxxxxxxxxxxxxxxxxxxxxx00002")
		return err
	}

	//预约人数已满
	if basicCfg.Capacity <= capUsed {
		return fmt.Errorf(ErrAppointmentString)
	}

	//通过planid 找到checkupcodes
	if app.PlanId != "" {
		if err := appPlan(tx, app, basicCfg); err != nil {
			fmt.Println("xxxxxxxxxxxxxxxxxxxxxx00003")
			return err
		}
	}

	fmt.Printf("--------------> app %v\n", app.SaleCodes)
	app.AppointedNum = genAppointNum(capUsed+1, basicCfg.AvoidNumbers)

	return saveAppoint(tx, app)
}

func deleteAppointment(tx *sql.Tx, a *Appointment) (err error) {
	sqlStr := ""
	//调整 TABLE_SaleRecords
	var salecodes, checkups []string
	var checkupsUsed map[string]int
	appointdatestring := time.Unix(a.AppointTime, 0).Format("2006-01-02")
	if salecodes, err = GetSalesByPlanID(tx, a.PlanId); err != nil {
		glog.Errorln("CancelAPpointment GetSalesByplan err", err.Error())
		return
	}
	if basic, err := org.GetConfigBasic(a.OrgCode); err != nil {
		glog.Errorln("appoint.deleteAppointment GetConfigBasic err", err.Error())
		return err
	} else {
		if checkups, err = getCheckupsBySales(basic.IpAddress, salecodes); err != nil {
			return err
		}
	}

	if checkupsUsed, err = GetCheckupsUsed(tx, a.OrgCode, appointdatestring, checkups); err != nil {
		glog.Errorln("CancelAPpointment GetSalesUsed err", err.Error())
		return
	}

	for sale_code, saleUsed := range checkupsUsed {
		sqlStr = SetCheckupAppointed_SQL(saleUsed-1, a.OrgCode, appointdatestring, sale_code)
		if _, err = tx.Exec(sqlStr); err != nil {
			fmt.Println("CancelAPpointment SetItemAppointed_SQL err", err.Error())
			return
		}
	}

	//capUserd, err := getAppDateRecordNum(tx, a.OrgCode, appointdatestring)
	//if err != nil {
	//	return
	//}

	return
}

//根据organization code , checkup code 获取分院特殊项目的限制
func GetCheckupLimit(tx *sql.Tx, orgCode string, ckCodes []string) (map[string]int, error) {
	ckStr := make([]string, len(ckCodes))
	for id, code := range ckCodes {
		ckStr[id] = fmt.Sprintf(`'%s'`, code)
	}

	sql := fmt.Sprintf("SELECT capacity,checkup_code FROM %s  WHERE org_code = '%s' AND checkup_code IN (%s)",
		org.TABLE_ORG_CON_SPECIAL, orgCode, strings.Join(ckStr, ","))

	rows, err := tx.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]int)
	var limit int
	var code string
	for rows.Next() {
		if err = rows.Scan(&limit, &code); err != nil {
			return nil, err
		}
		result[code] = limit
	}

	return result, nil
}

//得到分院某日的ｃｈｅｃｋｕｐ使用情况
func GetCheckupsUsed(tx *sql.Tx, orgCode, date string, checkups []string) (map[string]int, error) {
	checkupStr := make([]string, len(checkups))
	for k, code := range checkups {
		checkupStr[k] = fmt.Sprintf(`'%s'`, code)
	}

	rows, err := tx.Query(`SELECT DISTINCT checkup_code, count(checkup_code) OVER (PARTITION BY checkup_code)
	 FROM go_appoint_checkup_records where not cancel AND date = $1 AND org_code = $2 AND checkup_code in ($3)`, date, orgCode, pq.Array(checkups))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]int)

	var count int
	var checkup_code string
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&checkup_code, &count); err != nil {
			return nil, err
		}
		result[checkup_code] = count
	}

	return result, nil
}

func SetCheckupAppointed_SQL(used int, orgcode, date, sale string) string {
	var sql string
	if used == 0 {
		sql = fmt.Sprintf("DELETE FROM %s WHERE org_code = '%s' AND checkup_code = '%s' AND date = '%s' ",
			T_CHECKUP_RECORD, orgcode, sale, date)
		return sql
	}

	if used == 1 { //insert
		sql = fmt.Sprintf("INSERT INTO %s (org_code,date,used,checkup_code) VALUES ('%s','%s','%d','%s')",
			T_CHECKUP_RECORD, orgcode, date, used, sale)
		return sql
	}
	sql = fmt.Sprintf("UPDATE %s SET used = %d WHERE  org_code = '%s' AND date = '%s' AND checkup_code = '%s' ",
		T_CHECKUP_RECORD, used, orgcode, date, sale)

	return sql
}

func genAppointNum(num int, sortedNoUsedNums []int64) int {
	for _, noUsed := range sortedNoUsedNums {
		if int(noUsed) == num {
			return genAppointNum(num+1, sortedNoUsedNums)
		}
	}

	return num
}

func GetApp_for_wechatsByAppointments(a []Appointment) []App_For_WeChat {
	var afws []App_For_WeChat
	var afw App_For_WeChat
	for _, v := range a {
		afw.Name = v.Operator
		afw.AppointDate = time.Unix(v.AppointTime, 0).Format("2006-01-02")
		afw.PlanId = v.PlanId
		afw.Org_code = v.OrgCode

		if org, err := org.GetOrgByCode(afw.Org_code); err != nil {
			continue
		} else {
			afw.Org_Name = org.Name
		}
		afw.OperateTime = time.Unix(v.OperateTime, 0).Format("2006-01-02")
		afw.Serve_Mobile = "400400"
		afw.Status = v.Status
		if v.IfCancel {
			afw.Status = "已取消"
		}
		afw.AppID = v.ID
		afw.Reportid = v.ReportId
		afws = append(afws, afw)
	}
	return afws
}

func saveAppoint(tx *sql.Tx, app *Appointment) (err error) {
	sql := fmt.Sprintf(`INSERT INTO %s(id, appointtime, appointdate, org_code, planid, cardtype, cardno,
	mobile, appointor, appointorid, merrystatus, status, appoint_channel, company, "group", remark, operator,
	operatetime, orderid, commentid, appointednum, reportid, ifsingle, ifcancel, sale_codes)
	VALUES ('%s', '%d', '%s', '%s', '%s', '%s', '%s',
	'%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s',
	'%d', '%s', '%s', '%d', '%s', '%v', '%v' ,$1)
	ON CONFLICT (id) DO UPDATE SET
	appointtime=EXCLUDED.appointtime, org_code=EXCLUDED.org_code, planid=EXCLUDED.planid,
	cardtype=EXCLUDED.cardtype, cardno=EXCLUDED.cardno, mobile=EXCLUDED.mobile, appointor=EXCLUDED.appointor,
	merrystatus=EXCLUDED.merrystatus, status=EXCLUDED.status, appoint_channel=EXCLUDED.appoint_channel,
	company=EXCLUDED.company, "group"=EXCLUDED."group", remark=EXCLUDED.remark,
	operator=EXCLUDED.operatetime=EXCLUDED.operatetime, ifsingle=EXCLUDED.ifsingle, ifcancel=EXCLUDED.ifcancel`,

		T_APPOINTMENT, app.ID, app.AppointTime, app.AppointDate, app.OrgCode, app.PlanId, app.CardType,
		app.CardNo, app.Mobile, app.Appointor, app.Appointorid, app.MerryStatus, app.Status,
		app.Appoint_Channel, app.Company, app.Group, app.Remark, app.Operator, app.OperateTime,
		app.OrderID, app.CommentID, app.AppointedNum, app.ReportId, app.IfSingle, app.IfCancel)

	_, err = tx.Exec(sql, pq.Array(app.SaleCodes))
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxx00004")
	return
}
