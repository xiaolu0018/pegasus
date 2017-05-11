package sms

import (
	"fmt"
	"strconv"
	"strings"

	"encoding/json"

	"bjdaos/pegasus/pkg/common/util/methods"
)

//todo serveip 需可配置

const MessageHead = "【迪安诊断】"

func SendMessage(cellphoneNums []string, message string) ([]byte, int, error) {

	if !strings.Contains(message, MessageHead) {
		return nil, 0, fmt.Errorf("发送短信必须有 %s", MessageHead)
	}

	sms, err := NewSMS()
	if err != nil {
		return nil, 0, err
	}
	sms.Phones = strings.Join(cellphoneNums, ",")
	sms.Content = message
	return methods.Go_Through_HttpWithBody("POST", "http://apis.hzfacaiyu.com", "/sms/openCard", "", sms)
}

func SetMessageCheckNum(num int) string {
	return fmt.Sprintf("【迪安诊断】您的预约验证码为：%d", num)
}

func getSign(s SMS) (string, error) {
	signMap := map[string]string{}
	signMap["userName"] = s.UserName
	signMap["userPassword"] = s.UserPassword
	signMap["tradeNo"] = s.TradeNo
	marshal, err := json.Marshal(signMap)
	if err != nil {
		return "", err
	}
	a := AesEncrypt{}
	sign, err := a.EncryptByCBCNoPadding(marshal)
	var result string
	for _, s := range sign {
		B16 := strconv.FormatInt(int64(s), 16)
		if len(B16) == 1 {
			B16 = "0" + B16
		}
		result += strings.ToUpper(B16)
	}
	return result, err
}
