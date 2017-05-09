package user

import (
	"fmt"
	"strings"

	"github.com/golang/glog"

	"bjdaos/pegasus/pkg/wc/common"
	"bjdaos/pegasus/pkg/wc/db"
)

func (u *User) Upsert() error {
	sqlStr := fmt.Sprintf("INSERT INTO %s(id,openid,cardtype,cardno,mobile,name,sex,merrystatus,address_province,address_city,address_district,address_details,ifonlyneed_electronic_report,healthid) "+
		"VALUES('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%v','%s') ON CONFLICT (id) DO UPDATE SET cardtype=EXCLUDED.cardtype,cardno=EXCLUDED.cardno,"+
		"mobile=EXCLUDED.mobile,name=EXCLUDED.name,sex=EXCLUDED.sex,merrystatus=EXCLUDED.merrystatus,address_province=EXCLUDED.address_province,address_city=EXCLUDED.address_city,"+
		"address_district=EXCLUDED.address_district,address_details=EXCLUDED.address_details,ifonlyneed_electronic_report=EXCLUDED.ifonlyneed_electronic_report",
		TABLE_USER, u.ID, u.OpenID, u.CardType, u.CardNo, u.Mobile, u.Name, u.Sex, u.IsMarry, u.Address.Province, u.Address.City, u.Address.District,
		u.Address.Details, u.IsDianziReport, "")

	if _, err := db.GetDB().Exec(sqlStr); err != nil {
		return err
	}
	return nil
}

func GetUsersByOpenids(openids []string) ([]User, error) {

	sqlStr := fmt.Sprintf("SELECT id,openid FROM %s WHERE openid IN(%s)", TABLE_USER, forSqlIn(openids))
	rows, err := db.GetDB().Query(sqlStr)
	if err != nil {
		glog.Errorln("user.GetUsersByOpenids Query err", err.Error())
		return nil, err
	}
	defer rows.Close()
	var id, openid string
	users := make([]User, 10)
	for rows.Next() {
		if err = rows.Scan(&id, &openid); err != nil {
			glog.Errorln("user.GetUsersByOpenids rows.Scan err", err.Error())
			return nil, err
		}
		users = append(users, User{ID: id, OpenID: openid})
	}
	if rows.Err() != nil {
		glog.Errorln("user.GetUsersByOpenids rows.Err err", rows.Err().Error())
		return nil, rows.Err()
	}
	return users, nil
}

func GetUserByid(id string) (*User, error) {
	sqlStr := fmt.Sprintf("SELECT openid,cardtype,cardno,mobile,name,sex,merrystatus,address_province,address_city,address_district,address_details,wc_nickname,wc_headimgurl,ifonlyneed_electronic_report FROM %s WHERE id = '%s'", TABLE_USER, id)
	var openid, cardtype, cardno, mobile, name, sex, merrystatus, address_province, address_city, address_district, address_details, wc_nickname, wc_headimgurl string
	var ifonlyneed_electronic_report bool
	err := db.GetDB().QueryRow(sqlStr).Scan(&openid, &cardtype, &cardno, &mobile, &name, &sex, &merrystatus, &address_province, &address_city, &address_district, &address_details, &wc_nickname, &wc_headimgurl, &ifonlyneed_electronic_report)
	if err != nil {
		glog.Errorln("user.GetUserByid rows.Scan err", err.Error())
		return nil, err
	}
	u := User{
		ID:       id,
		CardType: cardtype,
		CardNo:   cardno,
		Mobile:   mobile,
		Name:     name,
		IsMarry:  merrystatus,
		Address: common.Address{
			Province: address_province,
			City:     address_city,
			District: address_district,
			Details:  address_details,
		},
		IsDianziReport: ifonlyneed_electronic_report,
		Sex:            sex,
		WC_Info: WCUserInfo{
			NickName:    wc_nickname,
			Head_ImgUrl: wc_headimgurl,
		},
	}
	return &u, nil
}

//为了sql时添加查询添加in
func forSqlIn(arr []string) string {
	for k, v := range arr {
		arr[k] = fmt.Sprintf("'%s'", v)
	}
	return strings.Join(arr, ",")
}
