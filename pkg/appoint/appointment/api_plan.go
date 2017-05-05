package appointment

import (
	"fmt"

	"database/sql"
	"github.com/lib/pq"

	"bjdaos/pegasus/pkg/appoint/db"
)

func GetSaleCodesByplan(tx *sql.Tx, planid string) ([]string, error) {
	sql := fmt.Sprintf("SELECT sale_codes FROM %s WHERE id = '%s'", TABLE_PALN, planid)
	var itemStr pq.StringArray
	if err := tx.QueryRow(sql).Scan(&itemStr); err != nil {
		return nil, err
	}

	items := []string(itemStr)
	return items, nil
}

func GetPlanByID(planid string) (*Plan, error) {
	pl := Plan{}
	salecodes := pq.StringArray{}
	sql := fmt.Sprintf("SELECT id, name, avatar_img, detail_img, sale_codes, ifshow FROM %s WHERE id = '%s'", TABLE_PALN, planid)
	if err := db.GetDB().QueryRow(sql).Scan(&pl.ID, &pl.Name, &pl.AvatarImg, &pl.DetailImg, &salecodes, &pl.IfShow); err != nil {
		return nil, err
	}

	pl.SaleCodes = []string(salecodes)
	return &pl, nil
}

func GetPlans() ([]Plan, error) {
	ps := make([]Plan, 0)
	sqlStr := fmt.Sprintf("SELECT id,name,avatar_img,detail_img,sale_codes FROM %s", TABLE_PALN)
	rows, err := db.GetDB().Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	p := Plan{}
	salecodes := pq.StringArray{}
	for rows.Next() {
		if err = rows.Scan(&p.ID, &p.Name, &p.AvatarImg, &p.DetailImg, &salecodes); err != nil {
			return nil, err
		}
		p.SaleCodes = []string(salecodes)
		ps = append(ps, p)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return ps, nil
}
