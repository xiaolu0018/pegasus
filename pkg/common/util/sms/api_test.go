package sms

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSendMessage(t *testing.T) {
	_, _, err := SendMessage([]string{"18792442120"}, "【迪安诊断】您的预约验证码为：9999")
	if err != nil {
		t.Fatal(err)
	}
}

