package db

import (
	"database/sql"
	_ "github.com/lib/pq"

	"fmt"
	"path/filepath"
	"io/ioutil"
)

//"postgres://postgres:postgresql2016@192.168.199.216:5432/pinto?sslmode=disable"
var db *sql.DB
var err error

func initDB(addr string) error {
	db, err = sql.Open("postgres", addr)
	if err != nil {
		return err
	}
	return loadSQLFunction("pkg/reporter/sql/function.sql")
}

func Init(user, passwd, ip, port, db string) error {
	dbAddr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, passwd, ip, port, db)

	return initDB(dbAddr)
}

func GetDB() *sql.DB {
	return db
}

func loadSQLFunction(path string) (err error){
	path , err = filepath.Abs(path)

	var data []byte
	data, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}

	_, err = db.Exec(string(data))
	return
}