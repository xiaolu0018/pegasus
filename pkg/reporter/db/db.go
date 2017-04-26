package db

import (
	"database/sql"
	_ "github.com/lib/pq"

	"fmt"
	"path/filepath"
	"io/ioutil"
)

//"postgres://postgres:postgresql2016@192.168.199.216:5432/pinto?sslmode=disable"
var readDB *sql.DB
var err error

func Init(user, passwd, ip, port, db string) error {
	addr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, passwd, ip, port, db)

	readDB, err = sql.Open("postgres", addr)
	return err
}

func GetDB() *sql.DB {
	return readDB
}

func InitFunction(user, passwd, ip, port, db string) error {
	addr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, passwd, ip, port, db)

	writeDB, err := sql.Open("postgres", addr)
	if err != nil {
		return err
	}
	defer writeDB.Close()

	return loadSQLFunction(writeDB, "pkg/reporter/sql/function.sql")
}

func loadSQLFunction(db *sql.DB, path string) (err error){
	path , err = filepath.Abs(path)

	var data []byte
	data, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}

	_, err = db.Exec(string(data))
	return
}