package user

import (
	"bjdaos/pegasus/pkg/wc/db"
	"fmt"
	"github.com/golang/glog"
	//"github.com/lib/pq"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"testing"
)

func dbinit() {
	if err := db.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto"); err != nil {
		fmt.Println("dbinit", err)
	}
}
func TestHealth_Upsert(t *testing.T) {
	dbinit()
	h := Health{
		Id:           bson.NewObjectId().Hex(),
		Past_history: []string{"心脑血管", "半身不随", "阿尔兹海默症"},
		//Past_history:[]string{},
		//Psychological_pressure:[]string{},
	}
	//sqlStr := fmt.Sprintf("INSERT INTO %s (id,past_history,family_medical_history,exam_frequency,past_exam_exception,psychological_pressure,food_habits,eating_habits,drink_habits,smoke_history)VALUES('%v','%s','%v','%v','%v','%v','%v','%v','%v','%v') ON CONFLICT (id) DO UPDATE SET id=EXCLUDED.id,past_history=EXCLUDED.past_history,family_medical_history=EXCLUDED.family_medical_history,exam_frequency=EXCLUDED.exam_frequency,past_exam_exception=EXCLUDED.past_exam_exception,psychological_pressure=EXCLUDED.psychological_pressure,food_habits=EXCLUDED.food_habits,eating_habits=EXCLUDED.eating_habits,drink_habits=EXCLUDED.drink_habits,smoke_history=EXCLUDED.smoke_history",
	//	TABLE_HEALTH, h.Id, pq.Array(h.Past_history), pq.Array(h.Family_medical_history), pq.Array(h.Exam_frequency), pq.Array(h.Past_exam_exception), pq.Array(h.Psychological_pressure), pq.Array(h.Food_habits),
	//	pq.Array(h.Eating_habits), pq.Array(h.Drink_habits), pq.Array(h.Smoke_history))
	sqlStr := fmt.Sprintf("INSERT INTO %s (id,past_history)VALUES('%v',%s)",
		TABLE_HEALTH, h.Id, GetArraySqlString(h.Past_history))
	fmt.Println("sqlstr", sqlStr)
	if rs, err := db.GetDB().Exec(sqlStr); err != nil {
		glog.Errorln("user.Upsert TABLE_HEALTH Exec err", err.Error())
		return
	} else {
		fmt.Println("rs", rs)
	}

}

func GetArraySqlString(s []string) string {
	var ss []string
	for _, v := range s {
		ss = append(ss, fmt.Sprintf(`"%s"`, v))
	}
	return `'{` + strings.Join(ss, ",") + `}'`
}
