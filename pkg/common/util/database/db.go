package database

import (
	"database/sql"
	_ "github.com/lib/pq"

	"fmt"
)

//"postgres://postgres:postgresql2016@192.168.199.216:5432/pinto?sslmode=disable"
func Init(user, passwd, ip, port, db string) (*sql.DB, error) {
	addr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, passwd, ip, port, db)

	return sql.Open("postgres", addr)
}
