package appointment

import "database/sql"

func addAppCheckupRecords(tx *sql.Tx, appointID, orgCode, date string, checkupCodes []string) error {
	for i := range checkupCodes {
		if _, err := tx.Exec(`INSERT INTO `+T_CHECKUP_RECORD+
			`(appoint_id, checkup_code, org_code, date) VALUES($1, $2, $3, $4)`,
			appointID, checkupCodes[i], orgCode, date); err != nil {
			return err
		}
	}

	return nil
}

func getAppDateRecordNum(tx *sql.Tx, orgCode, date string) (cap int, err error) {
	err = tx.QueryRow(`SELECT COUNT(*) FROM `+T_APPOINTMENT+
		` WHERE org_code = $1 AND appointdate = $2 AND NOT ifcancel`, orgCode, date).Scan(&cap)
	return
}
