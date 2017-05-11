package sms

import (
	"bjdaos/pegasus/pkg/common/util/md5"
	"github.com/1851616111/util/rand"
	"strconv"
	"time"
)

type SMS struct {
	Sign         string `json:"sign"`
	UserName     string `json:"userName"`
	Phones       string `json:"phones"`
	Content      string `json:"content"`
	UserPassword string `json:"userPassword"`
	TradeNo      string `json:"tradeNo"`
	Etnumber     string `json:"etnumber"`
}

//todo username pw 应该可配置
func NewSMS() (*SMS, error) {
	sms := SMS{}

	username := "dianzd"
	pw := "YvDGbZ5d"

	sms.UserName = username
	sms.UserPassword = pw
	sms.TradeNo = time.Now().Format("200601021504999") + strconv.Itoa(rand.RandInt(100, 999))
	var err error
	if sms.Sign, err = getSign(sms); err != nil {
		return nil, err
	}

	sms.UserPassword = md5.Md5([]byte(pw))
	return &sms, nil
}
