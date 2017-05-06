package db

import (
	"database/sql"
	_ "github.com/lib/pq"

	"fmt"
)

//"postgres://postgres:postgresql2016@192.168.199.216:5432/pinto?sslmode=disable"
var readDB *sql.DB
var writeDB *sql.DB
var err error

func InitReadDB(user, passwd, ip, port, db string) error {
	addr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, passwd, ip, port, db)

	readDB, err = sql.Open("postgres", addr)
	return err
}

func InitWriteDB(user, passwd, ip, port, db string) error {
	addr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, passwd, ip, port, db)

	writeDB, err = sql.Open("postgres", addr)
	return err
}

func GetReadDB() *sql.DB {
	return readDB
}

func GetWriteDB() *sql.DB {
	return writeDB
}
