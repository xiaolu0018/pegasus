package menu

import (
	"github.com/1851616111/util/weichat/manager/test"
	"testing"
)

func TestCreateMenu(t *testing.T) {
	// url json转码的时候需要 setEscapeHTML防止&被转移
	appoint := `https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxd09c7682905819e6&redirect_uri=http%3a%2f%2fwww.elepick.com%2fapi&response_type=code&scope=snsapi_base&state=123#wechat_redirect`
	button := NewTopButton("迪安").AddSub(NewViewButton("预约体检", appoint)).AddSub(NewViewButton("我的报告", "http://www.elepick.com/api"))
	err := CreateMenu(button, Test.Dev_Basic_Token)
	if err != nil {
		t.Fatal(err)
	}
}
