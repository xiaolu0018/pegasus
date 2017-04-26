package user

import (
	"192.168.199.199/bjdaos/pegasus/pkg/wc/db"
	"database/sql"
	"fmt"
	"github.com/golang/glog"
	"github.com/lib/pq"
	"strconv"
	"strings"
	"time"
)

func (h *Health) Upsert(userid string) (err error) {

	var tx *sql.Tx
	tx, err = db.GetDB().Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	sqlStr := fmt.Sprintf("SELECT healthid FROM %s WHERE id='%s'", TABLE_USER, userid)
	var healthid string
	if err = tx.QueryRow(sqlStr).Scan(&healthid); err != nil {
		glog.Errorln("user.Upsert rows.Scan err", err.Error())
		return
	}
	h.Id = strconv.FormatInt(time.Now().UnixNano(), 10)
	if healthid != "" {
		h.Id = healthid
	}

	sqlStr = fmt.Sprintf("INSERT INTO %s (id,past_history,family_medical_history,exam_frequency,past_exam_exception,psychological_pressure,food_habits,eating_habits,drink_habits,smoke_history)VALUES('%v',%s,%v,%v,%v,%v,%v,%v,%v,%v) ON CONFLICT (id) DO UPDATE SET id=EXCLUDED.id,past_history=EXCLUDED.past_history,family_medical_history=EXCLUDED.family_medical_history,exam_frequency=EXCLUDED.exam_frequency,past_exam_exception=EXCLUDED.past_exam_exception,psychological_pressure=EXCLUDED.psychological_pressure,food_habits=EXCLUDED.food_habits,eating_habits=EXCLUDED.eating_habits,drink_habits=EXCLUDED.drink_habits,smoke_history=EXCLUDED.smoke_history",
		TABLE_HEALTH, h.Id, GetArraySqlString(h.Past_history), GetArraySqlString(h.Family_medical_history), GetArraySqlString(h.Exam_frequency), GetArraySqlString(h.Past_exam_exception),
		GetArraySqlString(h.Psychological_pressure), GetArraySqlString(h.Food_habits), GetArraySqlString(h.Eating_habits), GetArraySqlString(h.Drink_habits), GetArraySqlString(h.Smoke_history))
	if _, err = tx.Exec(sqlStr); err != nil {
		glog.Errorln("user.Upsert TABLE_HEALTH Exec err", err.Error())
		return
	}

	sqlStr = fmt.Sprintf("UPDATE %s SET healthid = '%s' WHERE id = '%s'", TABLE_USER, h.Id, userid)
	if _, err = tx.Exec(sqlStr); err != nil {
		glog.Errorln("user.Upsert TABLE_USER Exec err", err.Error())
		return
	}
	return nil
}

func GetHealth(userid string) (*Health, error) {

	var err error

	sqlStr := fmt.Sprintf("SELECT healthid FROM %s WHERE id='%s'", TABLE_USER, userid)
	var healthid string
	if err = db.GetDB().QueryRow(sqlStr).Scan(&healthid); err != nil {
		glog.Errorln("user.GetHealth 1 rows.Scan err", err.Error())
		return nil, err
	}

	sqlStr = fmt.Sprintf("SELECT id,past_history,family_medical_history,exam_frequency,past_exam_exception,psychological_pressure,food_habits,eating_habits,drink_habits,smoke_history "+
		"FROM %s WHERE id = '%s'", TABLE_HEALTH, healthid)
	var past_history, family_medical_history, exam_frequency, past_exam_exception, psychological_pressure, food_habits, eating_habits, drink_habits, smoke_history pq.StringArray
	var id string
	if err = db.GetDB().QueryRow(sqlStr).Scan(&id, &past_history, &family_medical_history, &exam_frequency, &past_exam_exception, &psychological_pressure,
		&food_habits, &eating_habits, &drink_habits, &smoke_history); err != nil {
		glog.Errorln("user.GetHealth 2 rows.Scan err", err.Error())
		return nil, err
	}
	h := Health{
		Id:                     id,
		Past_history:           []string(past_history),
		Family_medical_history: []string(family_medical_history),
		Exam_frequency:         []string(exam_frequency),
		Past_exam_exception:    []string(psychological_pressure),
		Psychological_pressure: []string(psychological_pressure),
		Food_habits:            []string(food_habits),
		Eating_habits:          []string(eating_habits),
		Drink_habits:           []string(drink_habits),
		Smoke_history:          []string(smoke_history),
	}
	return &h, nil
}

func GetArraySqlString(s []string) string {
	var ss []string
	for _, v := range s {
		ss = append(ss, fmt.Sprintf(`"%s"`, v))
	}
	return `'{` + strings.Join(ss, ",") + `}'`
}
