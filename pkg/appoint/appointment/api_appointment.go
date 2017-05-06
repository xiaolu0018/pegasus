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
	"bjdaos/pegasus/pkg/appoint/organization"
	"bjdaos/pegasus/pkg/common/util/methods"
	"errors"
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

	tx := &sql.Tx{}
	if tx, err = db.GetDB().Begin(); err != nil {
		return
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
	var errs error
	if a.PlanId != "" {
		errs = sendToPintoWithPlan(a)
	} else {
		errs = sendToPintoWithOutPlan(a)
	}

	if errs != nil {
		return err
	}
	return
}

func sendToPintoWithOutPlan(a *Appointment) error {
	result, statuscode, errs := methods.Go_Through_HttpWithBody("POST", "http://192.168.199.198:9300", "/api/create/bookrecord", "", a)
	if errs != nil || statuscode != 200 {
		for i := 0; i < 3; i++ {
			if result, statuscode, errs = methods.Go_Through_HttpWithBody("POST", "http://192.168.199.198:9300", "/api/create/bookrecord", "", a); errs == nil && statuscode == 200 {
				break
			}
		}
	}
	resultmap := map[string]string{}

	json.Unmarshal(result, &resultmap)

	sqlStr := fmt.Sprintf("UPDATE %s SET bookno = '%s' WHERE id = '%s'", TABLE_APPOINTMENT, resultmap["bookno"], a.ID)
	if _, errs := db.GetDB().Exec(sqlStr); errs != nil {
		return errs
	}
	return nil
}

func sendToPintoWithPlan(a *Appointment) error {
	result, statuscode, errs := methods.Go_Through_HttpWithBody("POST", "http://192.168.199.198:9300", "/api/create/examwithplan", "", a)
	if errs != nil || statuscode != 200 {
		for i := 0; i < 3; i++ {
			if result, statuscode, errs = methods.Go_Through_HttpWithBody("POST", "http://192.168.199.198:9300", "/api/create/examwithplan", "", a); errs == nil && statuscode == 200 {
				break
			}
		}
	}
	resultmap := map[string]string{}

	json.Unmarshal(result, &resultmap)

	sqlStr := fmt.Sprintf("UPDATE %s SET bookno = '%s' WHERE id = '%s'", TABLE_APPOINTMENT, resultmap["bookno"], a.ID)
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
	sqlStr := ""
	sqlStr = fmt.Sprintf(`UPDATE %s SET ifcancel = 'true',status = '已取消' WHERE id = '%s'`, TABLE_APPOINTMENT, a.ID)
	if _, err = tx.Exec(sqlStr); err != nil {
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
	basic, err := organization.GetConfigBasic(a.OrgCode)
	if err != nil {
		return err
	}
	result := make(map[string]interface{})
	result["bookno"] = a.BookNo
	result["withplan"] = false
	if a.PlanId != "" {
		result["withplan"] = true
	}

	if _, statuscode, errs := methods.Go_Through_HttpWithBody("POST", basic.IpAddress, "/api/cancel/exam", "", result); errs == nil && statuscode == 200 {
		return nil
	} else {
		for i := 0; i < 3; i++ {
			if _, statuscode, errs = methods.Go_Through_HttpWithBody("POST", basic.IpAddress, "/api/cancel/exam", "", result); errs == nil && statuscode == 200 {
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
		`company,"group",remark,operator,operatetime,orderid,commentid,appointednum,ifsingle,ifcancel,sales_code,bookno FROM %s WHERE id = '%s'`, TABLE_APPOINTMENT, appointid)

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
	sqlStr := fmt.Sprintf("SELECT count(*) FROM %s WHERE cardno = '%s' AND appointtime = %d", TABLE_APPOINTMENT, cardno, appointtime)
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
		TABLE_APPOINTMENT, beginTimeSql, search, org_code, endTimeSql, userid)).Scan(&totalnums); err != nil {
		glog.Errorf("GetQueryAll: sql  err %v\n", err)
		return nil, 0, err
	}
	sqlStr := fmt.Sprintf("SELECT id,appointtime,org_code,planid,cardtype,cardno,mobile,appointor,appointorid,merrystatus,status,appoint_channel,"+
		`company,"group",remark,operator,operatetime,orderid,commentid,appointednum,reportid,ifsingle,ifcancel FROM %s WHERE  %s %s %s %s %s LIMIT '%d' OFFSET '%d' `,
		TABLE_APPOINTMENT, beginTimeSql, search, org_code, endTimeSql, userid, page_size, page_index)

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
		app.AppointTimeString = time.Unix(app.AppointTime, 0).Format("2006-01-02")
		app.OperateTimeString = time.Unix(app.OperateTime, 0).Format("2006-01-02")
		apps[k] = app
	}

	return apps, totalnums, nil
}

func addAppointment(tx *sql.Tx, a *Appointment) (err error) {
	itemLimits := make(map[string]int)
	itemused := make(map[string]int)
	var sales []string
	var sqlStr string
	if a.PlanId != "" {
		if sales, err = GetSaleCodesByplan(tx, a.PlanId); err != nil {
			glog.Errorf("planid err", err.Error())
			return
		}
	}

	if len(sales) == 0 {
		return errors.New("plan sales empty")
	}



	date := time.Unix(a.AppointTime, 0).Format("2006-01-02")
	if itemLimits, err = GetLimit(tx, a.OrgCode, sales); err != nil {
		glog.Errorf("GetLimit err", err.Error())
		return
	}

	if itemused, err = GetSalesUsed(tx, a.OrgCode, date, sales); err != nil {
		glog.Errorf("GetItemAppointedNum err ", err.Error())
		return
	}

	for item, itemLimit := range itemLimits {
		if appedCount, ok := itemused[item]; ok {
			if itemLimit <= appedCount {
				return fmt.Errorf(ErrAppointmentString)
			}
		}
	}
	fmt.Println("itemAppoint", itemused)
	//更新 该分院的 特殊项目的预约数量
	for item, used := range itemused {
		sqlStr = SetSaleAppointed_SQL(used+1, a.OrgCode, date, item)
		fmt.Println("sqlstr", sqlStr, used+1)
		if _, err = tx.Exec(sqlStr); err != nil {
			glog.Errorf("SetItemAppointed_SQL", err.Error())
			return
		}
	}

	var capacity, warnnum, capacityUsed int
	var avoidNumbers []int64
	if capacity, warnnum, avoidNumbers, err = GetOrg_Config(tx, a.OrgCode); err != nil {
		glog.Errorf("GetOrg_Config___", err.Error())
		return
	}

	if capacityUsed, err = GetCapacityUsed(tx, a.OrgCode, date); err != nil {
		glog.Errorf("GetCapacityused___", capacityUsed)
		return
	}
	//预约人数已满
	if capacity <= capacityUsed {
		return fmt.Errorf(ErrAppointmentString)
	}

	if (capacityUsed + 1) == warnnum {
		//todo 提醒相关人员
	}

	//更新该分院的总容量
	sqlStr = SetCapacityUsed_SQL(a.OrgCode, date, capacityUsed+1)
	if _, err = tx.Exec(sqlStr); err != nil {
		glog.Errorf("SetCapacityUsed_SQL___", err.Error())
		return
	}

	a.AppointedNum = getAppointedNum((capacityUsed + 1), avoidNumbers)
	a.SaleCodes = sales
	//保存预约
	sqlStr = fmt.Sprintf("INSERT INTO %s(id,appointtime,org_code,planid,cardtype,cardno,mobile,appointor,appointorid,merrystatus,status,appoint_channel,"+
		`company,"group",remark,operator,operatetime,orderid,commentid,appointednum,reportid,ifsingle,ifcancel,sale_codes) `+
		"VALUES ('%s','%d','%s','%s','%s','%s','%s','%s','%s','%s','%s',"+
		"'%s','%s','%s','%s','%s','%d','%s','%s','%d','%s','%v','%v',$1) ON CONFLICT (id)DO UPDATE SET appointtime=EXCLUDED.appointtime, org_code=EXCLUDED.org_code , planid=EXCLUDED.planid"+
		`, cardtype=EXCLUDED.cardtype,cardno=EXCLUDED.cardno,mobile=EXCLUDED.mobile,appointor=EXCLUDED.appointor,merrystatus=EXCLUDED.merrystatus,status=EXCLUDED.status,appoint_channel=EXCLUDED.appoint_channel,`+
		`company=EXCLUDED.company,"group"=EXCLUDED."group",remark=EXCLUDED.remark,operator=EXCLUDED.operatetime=EXCLUDED.operatetime,ifsingle=EXCLUDED.ifsingle,ifcancel=EXCLUDED.ifcancel`,
		TABLE_APPOINTMENT, a.ID, a.AppointTime, a.OrgCode, a.PlanId, a.CardType, a.CardNo, a.Mobile, a.Appointor, a.Appointorid, a.MerryStatus, a.Status,
		a.Appoint_Channel, a.Company, a.Group, a.Remark, a.Operator, a.OperateTime, a.OrderID, a.CommentID, a.AppointedNum, a.ReportId, a.IfSingle, a.IfCancel)
	fmt.Println("sqlStr", sqlStr)
	if _, err = tx.Exec(sqlStr, pq.Array(a.SaleCodes)); err != nil {
		glog.Errorf("TABLE_AppointmentsqlStr", err.Error())
		return
	}
	return
}
func deleteAppointment(tx *sql.Tx, a *Appointment) (err error) {
	sqlStr := ""
	//调整 TABLE_SaleRecords
	var salecodes []string
	var salesUsed map[string]int
	appointdatestring := time.Unix(a.AppointTime, 0).Format("2006-01-02")
	if salecodes, err = GetSaleCodesByplan(tx, a.PlanId); err != nil {
		glog.Errorln("CancelAPpointment GetSalesByplan err", err.Error())
		return
	}

	if salesUsed, err = GetSalesUsed(tx, a.OrgCode, appointdatestring, salecodes); err != nil {
		glog.Errorln("CancelAPpointment GetSalesUsed err", err.Error())
		return
	}

	for sale_code, saleUsed := range salesUsed {
		sqlStr = SetSaleAppointed_SQL(saleUsed-1, a.OrgCode, appointdatestring, sale_code)
		if _, err = tx.Exec(sqlStr); err != nil {
			fmt.Println("CancelAPpointment SetItemAppointed_SQL err", err.Error())
			return
		}
	}

	//调整　TABLE_CapacityRecords
	var capacityused int
	if capacityused, err = GetCapacityUsed(tx, a.OrgCode, appointdatestring); err != nil {
		fmt.Println("CancelAPpointment GetCapacityUsed err", err.Error())
		return
	}
	sqlStr = SetCapacityUsed_SQL(a.OrgCode, appointdatestring, capacityused-1)
	if _, err = tx.Exec(sqlStr); err != nil {
		fmt.Println("CancelAPpointment SetCapacityUsed_SQL err", err.Error())
		return
	}
	return
}

//得到分院的每个特殊项目的人数限制
func GetLimit(tx *sql.Tx, orgcode string, itemcodes []string) (map[string]int, error) {

	itmeStr := make([]string, len(itemcodes))
	for id, itemcodes := range itemcodes {
		itmeStr[id] = fmt.Sprintf(`'%s'`, itemcodes)
	}

	var limit int
	var itemcode string
	sql := fmt.Sprintf("SELECT capacity,sale_code FROM %s  WHERE org_code = '%s' AND sale_code IN (%s)",
		organization.TABLE_ORG_CON_SPECIAL, orgcode, strings.Join(itmeStr, ","))

	rows, err := tx.Query(sql);
	if err != nil {
		return nil, err
	}

	result := make(map[string]int)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&limit, &itemcode); err != nil {
			return nil, err
		}
		result[itemcode] = limit
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return result, nil
}

func GetOrg_Config(tx *sql.Tx, orgcode string) (int, int, []int64, error) {
	sql_ := fmt.Sprintf("SELECT capacity,warnnum,avoidnumbers FROM %s  WHERE org_code = '%s'",
		organization.TABLE_ORG_CON_BASIC, orgcode)
	var capacity, warnnum int
	avoidNumbers := pq.Int64Array{}
	var err error

	rows := &sql.Rows{}
	if rows, err = tx.Query(sql_); err != nil {
		return 0, 0, nil, err
	}

	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&capacity, &warnnum, &avoidNumbers); err != nil {
			return 0, 0, nil, err
		}
	}
	if rows.Err() != nil {
		return 0, 0, nil, rows.Err()
	}

	return capacity, warnnum, []int64(avoidNumbers), nil
}

func GetCapacityUsed(tx *sql.Tx, orgcode, date string) (int, error) {
	sql_ := fmt.Sprintf("SELECT used FROM %s WHERE org_code = '%s' AND date = '%s'",
		TABLE_CapacityRecords, orgcode, date)
	var err error
	var count int
	rows := &sql.Rows{}
	if rows, err = tx.Query(sql_); err != nil {
		return 0, err
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return 0, err
		}
	}
	if rows.Err() != nil {
		return 0, rows.Err()
	}

	return count, nil
}

func SetCapacityUsed_SQL(orgcode, date string, count int) string {
	sql := ""
	if count == 0 {
		sql = fmt.Sprintf("DELETE FROM %s WHERE org_code = '%s' AND date = '%s' ",
			TABLE_CapacityRecords, orgcode, date)
		return sql
	}

	if count == 1 { //insert
		sql = fmt.Sprintf("INSERT INTO %s (org_code,date,used) VALUES ('%s','%s','%d')",
			TABLE_CapacityRecords, orgcode, date, count)
		return sql
	}
	sql = fmt.Sprintf("UPDATE %s SET used = %d WHERE  org_code = '%s' AND date = '%s'",
		TABLE_CapacityRecords, count, orgcode, date)

	return sql
}

//得到分院的已经预约过的特殊项目及其预约人数
func GetSalesUsed(tx *sql.Tx, orgcode, date string, sales []string) (map[string]int, error) {
	itmeStr := make([]string, len(sales))
	for k, salecode := range sales {
		itmeStr[k] = fmt.Sprintf(`'%s'`, salecode)
	}

	//todo 这里应该进一步通过sales查询checkup
	sql_ := fmt.Sprintf("SELECT used,sale_code FROM %s  WHERE org_code = '%s' AND date = '%s' AND sale_code IN (%s)",
		TABLE_SaleRecords, orgcode, date, strings.Join(itmeStr, ","))
	var used int
	var sale_code string
	var rows *sql.Rows
	var err error
	if rows, err = tx.Query(sql_); err != nil {
		return nil, err
	}
	result := make(map[string]int)

	//初始化预约的项目，以便后面判断是更新还是插入
	for _, item := range sales {
		result[item] = 0
	}

	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&used, &sale_code); err != nil {
			return nil, err
		}
		result[sale_code] = used
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return result, nil
}

func SetSaleAppointed_SQL(used int, orgcode, date, sale string) string {

	sql := ""
	if used == 0 {
		sql = fmt.Sprintf("DELETE FROM %s WHERE org_code = '%s' AND sale_code = '%s' AND date = '%s' ",
			TABLE_SaleRecords, orgcode, sale, date)
		return sql
	}

	if used == 1 { //insert
		sql = fmt.Sprintf("INSERT INTO %s (org_code,date,used,sale_code) VALUES ('%s','%s','%d','%s')",
			TABLE_SaleRecords, orgcode, date, used, sale)
		return sql
	}
	sql = fmt.Sprintf("UPDATE %s SET used = %d WHERE  org_code = '%s' AND date = '%s' AND sale_code = '%s' ",
		TABLE_SaleRecords, used, orgcode, date, sale)

	return sql
}

func getAppointedNum(n int, sortArray []int64) int {
	i := 0
	for k, v := range sortArray {
		if int64(n) < v {
			i = k
			break
		}
	}
	return n + i
}

func GetApp_for_wechatsByAppointments(a []Appointment) []App_For_WeChat {
	var afws []App_For_WeChat
	var afw App_For_WeChat
	for _, v := range a {
		afw.Name = v.Operator
		afw.AppointDate = time.Unix(v.AppointTime, 0).Format("2006-01-02")
		afw.PlanId = v.PlanId
		afw.Org_code = v.OrgCode

		if org, err := organization.GetOrgByCode(afw.Org_code); err != nil {
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
		afw.Appid = v.ID
		afw.Reportid = v.ReportId
		afws = append(afws, afw)
	}
	return afws
}
