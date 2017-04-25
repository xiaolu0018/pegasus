package model

import (
	"192.168.199.199/bjdaos/pegasus/pkg/reporter/db"
	"192.168.199.199/bjdaos/totem/pkg/types"
	"fmt"

	"testing"
)

func Test(t *testing.T) {
	//name := "michael"
	//
	//name = fmt.Sprintf(`%%%s%%`, name)
	//fmt.Println(name)	iRet := GetPagesNumLimitPageNo(page, totalnums)
	if err := db.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto"); err != nil {
		fmt.Println("dbinit", err)
	}
	var rowsdate []types.PrintInfo
	rows, err := db.GetDB().Query(fmt.Sprintf(`SELECT e.examination_no, e.is_group, e.enterprise_name, e.checkupdate, p.name, p.sex, p.card_no, e.status
		FROM examination AS e , person AS p WHERE e.person_code = p.person_code and e.hos_code= '%s'
		%s AND status = %d  %s %s %s %s
		order by e.checkupdate desc
		limit '%d' offset '%d'`, "0001002", "", 1080, "", "", "", "", 20, 0))
	if err != nil {
		fmt.Errorf("GetQueryAll: sql return err %v\n", err)
	}
	for rows.Next() {
		var row = new(types.PrintInfo)
		if err := rows.Scan(&row.Ex_No, &row.Group, &row.Enterprise, &row.Ex_CkDate, &row.Name, &row.Sex, &row.CardNo, &row.Status); err != nil {
			fmt.Errorf("GetQueryAll: sql rows.Scan%v\n", err)
		}
		rowsdate = append(rowsdate, *row)
	}
	fmt.Println("rowsdate", len(rowsdate), rowsdate)

	if rows.Err() != nil {
		fmt.Errorf("GetQueryAll: sql rows.Err() %v\n", err)
	}

}
