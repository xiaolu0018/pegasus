package pinto

import (
	"bjdaos/pegasus/pkg/common/types"
	"database/sql"
	"fmt"
	"strings"
)

func ListCheckups(db *sql.DB) ([]types.Checkup, error) {
	rows, err := db.Query(`SELECT checkup_code, checkup_name, brief_name
	FROM checkup order by order_position`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	l := []types.Checkup{}
	var ck types.Checkup
	for rows.Next() {
		ck = types.Checkup{}
		if err = rows.Scan(&ck.Checkup_code, &ck.Checkup_name, &ck.Brief_name); err == nil {
			l = append(l, ck)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return l, nil
}

//SELECT * FROM examination WHERE examination_no = '0001001170003658';
//SELECT * FROM examination_charge_total WHERE examination_no = '0001001170003658';
//SELECT * FROM examination_log WHERE examination_no = '0001001170003658';
//SELECT * FROM examination_sale WHERE examination_no = '0001001170003658';

func GetCheckupBySaleCode(db *sql.DB, code string) ([]types.Checkup, error) {
	rows, err := db.Query(`select c.checkup_code, c.checkup_name, c.brief_name
	from sale_checkup sc, checkup c
	where sc.checkup_code = c.checkup_code
	and sc.sale_code = $1`, code)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	l := []types.Checkup{}
	var ck types.Checkup
	for rows.Next() {
		ck = types.Checkup{}
		if err = rows.Scan(&ck.Checkup_code, &ck.Checkup_name, &ck.Brief_name); err == nil {
			l = append(l, ck)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return l, nil

}

func GetCheckupCodesBySaleCodes(db *sql.DB, salecodes []string) ([]string, error) {

	itmeStr := make([]string, len(salecodes))
	for k, salecode := range salecodes {
		itmeStr[k] = fmt.Sprintf(`'%s'`, salecode)
	}
	sqlStr := fmt.Sprintf("SELECT checkup_code FROM sale_checkup WHERE sale_code IN (%s)", strings.Join(itmeStr, ","))
	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var checkupcode string
	var checkupcodes []string
	for rows.Next() {
		if err = rows.Scan(&checkupcode); err != nil {
			return nil, err
		}
		checkupcodes = append(checkupcodes, checkupcode)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return checkupcodes, nil
}
