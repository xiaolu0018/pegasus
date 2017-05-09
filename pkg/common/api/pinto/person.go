package pinto

import (
	"bjdaos/pegasus/pkg/common/types"
	"database/sql"
	"fmt"
	"bjdaos/pegasus/pkg/common/util/timeutil"
)

func CreatePerson(db *sql.DB, person types.Person) error {
	sqlStr := fmt.Sprintf("INSERT INTO person(sex,card_no,is_marry,name,cellphone,createtime,person_code,idcard_type_code,hos_code)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)")
	_, err := db.Exec(sqlStr, person.Sex, person.CardNo, person.IsMarry, person.Name, person.CellPhone, person.CreateTime, person.PersonCode, person.IdcardTypeCode, person.HosCode)
	return err
}

func FilterPersonByAppoint(a *Appointment)*types.Person{
	var p types.Person
	p.CreateTime = a.TimeNow.Format(timeutil.FROMAT_DAY)
	p.Sex = SexToCode[a.Sex]
	p.CardNo = a.CardNo
	p.IdcardTypeCode = IdCardToCode[a.CardType]
	p.CellPhone = a.Mobile
	p.HosCode = a.OrgCode
	p.IsMarry = MarryToCode[a.MerryStatus]
	p.Name = a.Appointor
	p.PersonCode = a.TimeNow.Format(timeutil.FROMAT_SECOND)
	return &p
}
