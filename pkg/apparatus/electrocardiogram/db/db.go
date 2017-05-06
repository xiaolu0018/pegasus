package db


import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

//"postgres://postgres:postgresql2016@192.168.199.216:5432/pinto?sslmode=disable"
var db *sql.DB

func initDB(addr string) error {
	var err error
	db, err = sql.Open("postgres", addr)
	return err
}

func Init(user, passwd, ip, port, db string) error {
	dbAddr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, passwd, ip, port, db)
	return initDB(dbAddr)
}

func GetDB() *sql.DB {
	return db
}

