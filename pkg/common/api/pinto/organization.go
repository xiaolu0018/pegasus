package pinto

import (
	"database/sql"

	"192.168.199.199/bjdaos/pegasus/pkg/common/types"
)

func ListOrganizations(db *sql.DB) ([]types.Organization, error) {
	sql := `SELECT id, org_code, org_name FROM organization WHERE is_valid = 1 ORDER BY parent_code`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	l := []types.Organization{}
	var id, code, name string
	for rows.Next() {
		if err = rows.Scan(&id, &code, &name); err == nil {
			if len(code) == 7 {
				l = append(l, types.Organization{
					ID:   id,
					Code: code,
					Name: name,
				})
			}
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return l, nil
}
