package appointment

import (
	"fmt"
	"strings"

	"database/sql"
)

func GetSalesByplan(tx *sql.Tx, planid string) ([]string, error) {
	sql := fmt.Sprintf("SELECT sales FROM %s WHERE id = '%s'", TABLE_PALN, planid)
	var itemStr string
	if err := tx.QueryRow(sql).Scan(&itemStr); err != nil {
		return nil, err
	}

	items := strings.Split(itemStr[1:len(itemStr)-1], ",")
	return items, nil
}
