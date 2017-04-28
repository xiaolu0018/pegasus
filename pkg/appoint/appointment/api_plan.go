package appointment

import (
	"fmt"
	"strings"

	"database/sql"
	"github.com/lib/pq"

	"192.168.199.199/bjdaos/pegasus/pkg/appoint/db"
)

func GetCheckupsByplan(tx *sql.Tx, planid string) ([]string, error) {
	sql := fmt.Sprintf("SELECT checkups FROM %s WHERE id = '%s'", TABLE_PALN, planid)
	var itemStr string
	if err := tx.QueryRow(sql).Scan(&itemStr); err != nil {
		return nil, err
	}

	items := strings.Split(itemStr[1:len(itemStr)-1], ",")
	return items, nil
}

func GetPlanByID(planid string) (*Plan, error) {
	pl := Plan{}
	checkups := pq.StringArray{}
	sql := fmt.Sprintf("SELECT id, name, avatar_img, detail_img, checkups, ifshow FROM %s WHERE id = '%s'", TABLE_PALN, planid)
	if err := db.GetDB().QueryRow(sql).Scan(&pl.ID, &pl.Name, &pl.AvatarImg, &pl.DetailImg, &checkups, &pl.IfShow); err != nil {
		return nil, err
	}

	pl.Checkups = []string(checkups)
	return &pl, nil
}

func GetPlans() ([]Plan, error) {
	ps := make([]Plan, 0)
	sqlStr := fmt.Sprintf("SELECT id,name,avatar_img,detail_img,checkups FROM %s", TABLE_PALN)
	rows, err := db.GetDB().Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	p := Plan{}
	checkups := pq.StringArray{}
	for rows.Next() {
		if err = rows.Scan(&p.ID, &p.Name, &p.AvatarImg, &p.DetailImg, &checkups); err != nil {
			return nil, err
		}
		p.Checkups = checkups
		ps = append(ps, p)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return ps, nil
}
