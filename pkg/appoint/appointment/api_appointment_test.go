package appointment

import (
	"bjdaos/pegasus/pkg/appoint/db"
	"fmt"
	"strings"
	"testing"
	//"time"
	"bjdaos/pegasus/pkg/appoint/organization"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"reflect"
	"time"
)

func dbinit() {
	if err := db.Init("postgres", "postgres190@", "10.1.0.190", "5432", "pinto"); err != nil {
		fmt.Println("dbinit", err)
	}
}

func TestGetItemByplan(t *testing.T) {
	dbinit()
	planid := "007"
	sql := fmt.Sprintf("SELECT offdays FROM %s WHERE org_code = '%s'", "go_appoint_organization_basic_con", planid)
	var items string
	if err := db.GetDB().QueryRow(sql).Scan(&items); err != nil {
		fmt.Println(err)
	}
	//if is, ok := items.(string); ok {
	//	fmt.Println("fdfisis ", is)
	//}
	itemss := strings.Split(items[1:len(items)-1], ",")
	fmt.Println("fdf ", items, itemss)
}

func TestCreatAppoint(t *testing.T) {
	dbinit()
	a := Appointment{
		ID:              "",
		Appointor:       "name1",
		CardNo:          "cardid1",
		CardType:        "身份证",
		Mobile:          "18799552120",
		MerryStatus:     "未婚",
		Status:          "预约",
		Appoint_Channel: "微信",
		AppointedNum:    0,
		Sex:             "男",
		PlanId:          "1",
		OrgCode:         "0001002",
		AppointTime:     time.Now().Unix(),
		OperateTime:     time.Now().Unix(),
		OrderID:         "order13",
		Operator:        "operator13",
	}
	//
	err := a.CreateAppointment()
	fmt.Println("err", err)
}

func TestGetCapacityAppointedNum(t *testing.T) {
	dbinit()
	sqlstr := fmt.Sprintf("INSERT INTO %s (org_code,date,used,sale_code) VALUES ('%s','%s','%d','%s')",
		TABLE_SaleRecords, "000101", "2017-04-20", 2, "0002")
	result, err := db.GetDB().Exec(sqlstr)
	fmt.Println("re", result, err)
}

func TestGetAppointment(t *testing.T) {
	dbinit()
	a, err := GetAppointment("appoint13")
	fmt.Println(a, err)
}

func TestJson(t *testing.T) {
	a := Appointment{
		ID:              "112313dfdfdf",
		Appointor:       "httptext",
		CardNo:          "cardid1",
		CardType:        "cardType1",
		Mobile:          "mobile1",
		MerryStatus:     "未婚",
		Status:          "预约",
		Appoint_Channel: "微信",
		AppointedNum:    0,
		PlanId:          "1",
		OrgCode:         "000101",
		AppointTime:     time.Now().AddDate(0, 0, 2).Unix(),
		OperateTime:     time.Now().Unix(),
		OrderID:         "order13",
		Operator:        "operator13",
	}
	stirngddd := GetJsonType(a)
	fmt.Println("reture", stirngddd)
}

func TestJson__(t *testing.T) {
	//dbinit()
	//app, total, err := GetAppointmentList(0, 20, 0, 0, "", "")
	begintime := time.Time{}
	fmt.Println("time", GetDayFirstSec(begintime))

	a := Comment{}
	fmt.Println(GetJsonType(a))
}

func TestComment_Create(t *testing.T) {
	dbinit()
	c := Comment{}
	fmt.Println(c.Create("112313"))
}

func GetJsonType(strcut_ interface{}) string {
	tp := reflect.TypeOf(strcut_)

	v := reflect.ValueOf(strcut_)
	var key_vals []string
	var k_v string
	for i := 0; i < v.NumField(); i++ {
		key := tp.Field(i).Name
		val := v.Field(i)
		fmt.Println("key_%s,val_%s", key, val, v.Field(i).Type(), v.Field(i).Kind(), v.Field(i).String())
		if v.Field(i).Kind() == reflect.Int || v.Field(i).Kind() == reflect.Int64 || v.Field(i).Kind() == reflect.Float32 {
			k_v = fmt.Sprintf(`"%s":%v`, strings.ToLower(key), val)
		} else if v.Field(i).Kind() == reflect.Bool {
			k_v = fmt.Sprintf(`"%s":%v`, strings.ToLower(key), val)
		} else {
			k_v = fmt.Sprintf(`"%s":"%v"`, strings.ToLower(key), val)
		}

		key_vals = append(key_vals, k_v)

	}

	return `'{` + strings.Join(key_vals, ",") + `}'`
}

func GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

func TestChangeOrg(t *testing.T) {
	dbinit()
	sqlstr := fmt.Sprintf("UPDATE %s SET imageurl = '%s',detailsurl = '%s'", organization.TABLE_ORG, "", "")
	_, err := db.GetDB().Exec(sqlstr)
	fmt.Println("errr", err)
}
