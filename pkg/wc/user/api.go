package user

import (
	"192.168.199.199/bjdaos/pegasus/pkg/wc/db"
	"fmt"
)

func (u *User) Upsert() error {
	sqlStr := fmt.Sprintf("INSERT INTO %s(id,openid,cardtype,cardno,mobile,name,merrystatus,address_province,address_city,address_district,address_details,ifonlyneed_electronic_report,healthid) "+
		"VALUES('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s') ON CONFLICT (id) DO UPDATE SET cardtype=EXCLUDED.cardtype,cardno=EXCLUDED.cardno,"+
		"mobile=EXCLUDED.mobile,name=EXCLUDED.name,merrystatus=EXCLUDED.merrystatus,address_province=EXCLUDED.address_province,address_city=EXCLUDED.address_city,"+
		"address_city=EXCLUDED.address_city,address_district=EXCLUDED.address_district,address_details=EXCLUDED.address_details,ifonlyneed_electronic_report=EXCLUDED.ifonlyneed_electronic_report",
		TABLE_USER, u.ID, u.OpenID, u.CardType, u.CardNo, u.Mobile, u.Name, u.IsMarry, u.Address.Province, u.Address.City, u.Address.District,
		u.Address.Details, u.IsDianziReport, "")

	if _, err := db.GetDB().Exec(sqlStr); err != nil {
		return err
	}
	return nil
}

func (h *Health) Upsert() error {
	sqlStr := fmt.Sprintf("INSERT INTO %s (id,past_history,family_medical_history,exam_frequency,past_exam_exception,psychological_pressure,food_habits,eating_habits,drink_habits,smoke_history)VALUES('%v','%v','%v','%v','%v','%v','%v','%v','%v','%v') ON CONFLICT (id) DO UPDATE SET id=EXCLUDED.id,past_history=EXCLUDED.past_history,family_medical_history=EXCLUDED.family_medical_history,exam_frequency=EXCLUDED.exam_frequency,past_exam_exception=EXCLUDED.past_exam_exception,psychological_pressure=EXCLUDED.psychological_pressure,food_habits=EXCLUDED.food_habits,eating_habits=EXCLUDED.eating_habits,drink_habits=EXCLUDED.drink_habits,smoke_history=EXCLUDED.smoke_history",
		TABLE_HEALTH, h.Id, h.Past_history, h.Family_medical_history, h.Exam_frequency, h.Past_exam_exception, h.Psychological_pressure, h.Food_habits, h.Eating_habits, h.Drink_habits, h.Smoke_history)
	if _, err := db.GetDB().Exec(sqlStr); err != nil {
		return err
	}
	return nil
}
