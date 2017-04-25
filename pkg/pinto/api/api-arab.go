package api

import "database/sql"

var arabDB *sql.DB

func InitArabDB(db *sql.DB) error {
	pintoDB = db
	return nil
}

func GetOrders() {

}
