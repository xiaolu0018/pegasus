package api

import (
	"fmt"

	"database/sql"

	"192.168.199.199/bjdaos/pegasus/pkg/common/types"
)

var pintoDB *sql.DB

func InitPintoDB(db *sql.DB) error {
	pintoDB = db
	return nil
}

func ListAllOrgs() ([]types.Organization, error) {
	sql := `SELECT id, org_code, org_name FROM organization WHERE is_valid = 1 ORDER BY parent_code`
	rows, err := pintoDB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	l := []types.Organization{}
	var id, code, name string
	for rows.Next() {
		if err = rows.Scan(&id, &code, &name); err == nil {
			l = append(l, types.Organization{
				ID:   id,
				Code: code,
				Name: name,
			})
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return l, nil
}

func GetSalesByOrgCode(code string) ([]types.Sale, error) {
	sql := fmt.Sprintf(`SELECT sale_code, brief_name from sale where org_code = '%s' order by order_position`, code)

	rows, err := pintoDB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	l := []types.Sale{}
	var sale_code, brief_name string
	for rows.Next() {
		if err = rows.Scan(&sale_code, &brief_name); err == nil {
			l = append(l, types.Sale{
				Code:      sale_code,
				BriefName: brief_name,
			})
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return l, nil
}
