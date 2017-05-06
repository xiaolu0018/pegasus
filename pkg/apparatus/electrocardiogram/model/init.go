package model

import (
	"database/sql"
)

var DB *sql.DB

func Init(db *sql.DB) error {
	DB = db
	return nil
}
