package appointment

import (
	"fmt"
	"time"
	"bytes"
	"encoding/json"

	"github.com/golang/glog"
	"bjdaos/pegasus/pkg/appoint/db"
	"bjdaos/pegasus/pkg/common/util/methods"
	tm "bjdaos/pegasus/pkg/common/util/time"
)

//将体检状态 预约 改为体检中,或待评价 ，每5分钟更新一次
var ExamStatusToAppointStatus = map[int]string{
	1040: "体检中",
	1041: "体检中",
	1042: "体检中",
	1050: "待评价",
}

func StartController() {
	appBrokenTicker := time.NewTicker(time.Hour)
	appSyncTicker := time.NewTicker(time.Minute * 5)
	for {
		select {
		case <- appBrokenTicker.C:
			if err := appBrokenStatus(); err != nil {
				glog.Errorf("appoint controller: broken appoint status err %v\n", err)
			}
		case <- appSyncTicker.C:
			if err := appSyncStatus(); err != nil {
				glog.Errorf("appoint controller: sync appoint status err %v\n", err)
			}
		}

	}
}

//改变没有应约的状态
func appBrokenStatus() (err error) {
	sqlStr := fmt.Sprintf(`UPDATE %s SET status = '爽约' WHERE status = '%s'`, TABLE_APPOINTMENT, STATUS_SUCCESS)
	_, err = db.GetDB().Exec(sqlStr);
	return
}

func appSyncStatus() error {
		//只需查体检日期是当天的就行
		sqlStr := fmt.Sprintf(`SELECT a.bookno, a.org_code, basic.ip_address FROM %s a
		LEFT JOIN %s basic
		ON a.org_code = basic.org_code
		WHERE a.appointtime BETWEEN %d AND %d
		AND a.status IN ('%s','%s')`,
		TABLE_APPOINTMENT, TABLE_ORGANIZATION_CONFIG_BASIC, tm.TodayStartSec(time.Now()), tm.TodayEndSec(time.Now()), STATUS_SUCCESS, STATUS_EXAMING)

		rows, err := db.GetDB().Query(sqlStr)
		if err != nil {
			return err
		}
		var bookno, org_code, ip_address string

		org_booknos := map[string][]string{}
		ip_org := map[string][]string{}
		for rows.Next() {
			if err = rows.Scan(&bookno, &org_code, &ip_address); err != nil {
				glog.Errorln("appoint.changeAppoitmentStatusToExaming db.rows.Scan err ", err)
				rows.Close()
				continue
			}
			if _, ok := org_booknos[org_code]; ok {
				org_booknos[org_code] = append(org_booknos[org_code], bookno)
			} else {
				org_booknos[org_code] = []string{bookno}
			}

			if _, ok := ip_org[ip_address]; ok {
				ip_org[ip_address] = append(ip_org[ip_address], org_code)
			} else {
				ip_org[ip_address] = []string{org_code}
			}
		}
		rows.Close()

		for key, orgs := range ip_org {
			tmp_map := make(map[string]interface{})

			tmp_books := []string{}
			for _, org := range orgs {
				tmp_books = append(tmp_books, org_booknos[org]...)
			}

			tmp_map["bookno"] = tmp_books
			rebyte, _, err := methods.Go_Through_HttpWithBody("GET", key, "/api/exam/status", "", tmp_map)
			if err != nil {
				glog.Errorln("appoint.changeAppoitmentStatusToExaming Go_Through_HttpWithBody err ", err)
			}
			var result_https interface{}
			err = json.NewDecoder(bytes.NewReader(rebyte)).Decode(&result_https)

			if result_https == nil {
				continue
			}
			glog.Errorln("appoint.changeAppointmentStatusToExaming result_http___", result_https, err)

			for _, result_http := range result_https.([]interface{}) {
				if status, ok := result_http.(map[string]interface{})["status"]; ok && int(status.(float64)) > 0 {
					if bookno, ok := result_http.(map[string]interface{})["bookno"]; ok {
						appointstatus := ExamStatusToAppointStatus[int(status.(float64))]
						fmt.Println("appoint status", appointstatus)
						if appointstatus != "" {
							sqlStr = fmt.Sprintf("UPDATE %s SET status = '%s' WHERE bookno = '%s' AND status <> '%s'", TABLE_APPOINTMENT, appointstatus, bookno.(string), appointstatus)
							if _, err = db.GetDB().Exec(sqlStr); err != nil {
								glog.Errorln("appoint.changeAppointmentStatusToExaming UPDATE TABLE_APPOINTMENT err ", err)
							}
						}
					}
				}
			}
		}
	return nil
}
