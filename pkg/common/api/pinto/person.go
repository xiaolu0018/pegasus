package pinto

import (
	"192.168.199.199/bjdaos/pegasus/pkg/common/types"
	"database/sql"
	"fmt"
)

func CreatePerson(db *sql.DB, person types.Person) error {
	sqlStr := fmt.Sprintf("INSERT INTO person(sex,card_no,is_marry,name,cellphone,createtime,person_code,idcard_type_code,hos_code)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)")
	_, err := db.Exec(sqlStr, person.Sex, person.CardNo, person.IsMarry, person.Name, person.CellPhone, person.CreateTime, person.PersonCode, person.IdcardTypeCode, person.HosCode)
	return err
}
