package user

import (
	"bjdaos/pegasus/pkg/wc/db"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strings"
	"testing"
)

func TestUser_UpdateLabel(t *testing.T) {
	labelhealth := make(map[string][]string)
	labelhealth["bingshi"] = []string{"gaoxueya", "gaoxuezhi"}
	labelhealth["jiazushi"] = []string{"guanxinbing", "naogeng"}
	err := UpdateLabel(db.User(), bson.NewObjectId(), labelhealth)
	fmt.Println(err)
}

func TestGetUPSERTSQLStrByStruct(t *testing.T) {
	//h := Health{}
	//stringss := GetUPSERTSQLStrByStruct(h)
	us := User{}
	st := GetSelectString(us)
	fmt.Println("strgingsss", st)
}

func GetUPSERTSQLStrByStruct(struct_ interface{}) string {
	tp := reflect.TypeOf(struct_)
	v := reflect.ValueOf(struct_)
	key := make([]string, 0)
	values := make([]string, 0)
	EXCLUDED := make([]string, 0)
	for i := 0; i < tp.NumField(); i++ {
		key = append(key, strings.ToLower(tp.Field(i).Name))
		EXCLUDED = append(EXCLUDED, fmt.Sprintf("%s=EXCLUDED.%s", strings.ToLower(tp.Field(i).Name), strings.ToLower(tp.Field(i).Name)))
		if v.Field(i).Kind() == reflect.Int || v.Field(i).Kind() == reflect.Int64 || v.Field(i).Kind() == reflect.Float32 || v.Field(i).Kind() == reflect.Bool {
			values = append(values, "%v")
		} else {
			values = append(values, "'%v'")
		}
	}

	s1 := "INSERT INTO %S ("
	keystring := strings.Join(key, ",")
	s1 = s1 + keystring + ")"
	s2 := "VALUES(" + strings.Join(values, ",") + ") ON CONFLICT (id) DO UPDATE SET " + strings.Join(EXCLUDED, ",")

	s1 = s1 + s2
	return s1
}

func GetSelectString(s interface{}) string {
	tp := reflect.TypeOf(s)
	keys := make([]string, 0)
	for i := 0; i < tp.NumField(); i++ {
		keys = append(keys, strings.ToLower(tp.Field(i).Name))
	}
	return strings.Join(keys, ",")
}

func GetPtrForSql(s []string) string {
	arr := make([]string, 0, len(s))
	for _, v := range s {
		arr = append(arr, fmt.Sprintf("&%s", v))
	}
	return strings.Join(arr, ",")
}
