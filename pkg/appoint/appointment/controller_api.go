package appointment

import (
	"bjdaos/pegasus/pkg/appoint/db"
	tm "bjdaos/pegasus/pkg/common/util/time"
	"encoding/json"
	"fmt"
	"github.com/1851616111/util/http"
	"github.com/golang/glog"
	"time"
)

type org struct {
	orgCode    string
	ordAddress string
}

type OrgToBookOrders map[org][]string

func getTodayOrgBookOrders() (*OrgToBookOrders, error) {
	sqlStr := fmt.Sprintf(`SELECT a.bookno, a.org_code, basic.ip_address FROM %s a
		LEFT JOIN %s basic
		ON a.org_code = basic.org_code
		WHERE a.appointtime BETWEEN %d AND %d
		AND a.status IN ('%s','%s')`,
		TABLE_APPOINTMENT, TABLE_ORGANIZATION_CONFIG_BASIC, tm.TodayStartSec(time.Now()), tm.TodayEndSec(time.Now()), STATUS_SUCCESS, STATUS_EXAMING)

	rows, err := db.GetDB().Query(sqlStr)
	if err != nil {
		return nil, err
	}

	ret := OrgToBookOrders{}
	o := org{}
	var bookNo string
	for rows.Next() {
		if err := rows.Scan(&bookNo, &o.orgCode, &o.ordAddress); err != nil {
			glog.Errorf("appoint getTodayOrgBookOrders():　scan result err %v\n", err)
		}

		if _, exist := ret[o]; !exist {
			ret[o] = []string{}
		} else {
			ret[o] = append(ret[o], bookNo)
		}
	}

	return &ret, nil
}

func listOrgBookStatus(orgAddr string, bookNos []string) (map[string]int, error) {
	if len(bookNos) == 0 {
		return nil, nil
	}

	rsp, err := http.Send(&http.HttpSpec{
		URL:         orgAddr + "/api/exam/status",
		Method:      "GET",
		ContentType: http.ContentType_JSON,
		BodyParams:  http.NewBody().Add("booknos", bookNos),
	})

	if err != nil {
		return nil, err
	}

	result := map[string]int{}
	dc := json.NewDecoder(rsp.Body)
	dc.UseNumber()
	err = dc.Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func batchUpdateStatus(bookStatsM map[string]int) error {
	if bookStatsM == nil {
		return nil
	}
	for bookNo, status := range bookStatsM {
		if status != 0 {
			sql := fmt.Sprintf("UPDATE %s SET status = '%d' WHERE bookno = '%s' AND status <> '%d'", TABLE_APPOINTMENT, status, bookNo, status)
			if _, err := db.GetDB().Exec(sql); err != nil {
				return err
			}
		}
	}

	return nil
}
