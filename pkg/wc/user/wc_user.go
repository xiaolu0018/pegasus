package user

import (
	"encoding/json"
	"github.com/1851616111/util/http"

	"192.168.199.199/bjdaos/pegasus/pkg/wc/common"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/db"
	"errors"
)

type wcToken struct {
	openID      string
	accessToken string
}

type WCUserInfo struct {
	NickName    string `json:"nickname"`
	Head_ImgUrl string `json:"headimgurl"`
	Sex         int    `json:"sex"`
	common.Address
}

func (w *wcToken) updateUserWCInfo() error {
	i, err := getWCUserInfo(w.openID, w.accessToken)
	if err != nil {
		return err
	}

	return upsertUser(w.openID, i)
}

func getWCUserInfo(id, token string) (*WCUserInfo, error) {
	rsp, err := http.Send(&http.HttpSpec{
		URL:       "https://api.weixin.qq.com/sns/userinfo",
		Method:    "GET",
		URLParams: http.NewParams().Add("openid", id).Add("access_token", token).Add("lang", "zh_CN"),
	})
	if err != nil {
		return nil, err
	}

	ui := WCUserInfo{}
	if err := json.NewDecoder(rsp.Body).Decode(&ui); err != nil {
		return nil, err
	}

	return &ui, nil
}

func upsertUser(openid string, ui *WCUserInfo) error {
	var count int
	if err := db.GetDB().QueryRow(`select count(id) from `+TABLE_USER+` where openid =$1`, openid).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		return errors.New("user not found")
	}

	_, err := db.GetDB().Exec(`UPDATE `+TABLE_USER+`
	SET wc_nickname=$1, wc_sex=$2, wc_headimgurl=$3, wc_country=$4,wc_province=$5, wc_city=$6
	WHERE openid=$7`, ui.NickName, ui.Sex, ui.Head_ImgUrl, ui.Country, ui.Province, ui.City, openid)
	return err
}
