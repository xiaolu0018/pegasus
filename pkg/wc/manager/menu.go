package manager

import (
	"github.com/1851616111/util/weichat/manager/menu"
)

func CreateMenu(access_token string) error {
	appoint := `https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxd09c7682905819e6&redirect_uri=http%3a%2f%2fwww.elepick.com%2fapi%2fappoint&response_type=code&scope=snsapi_base&state=123#wechat_redirect`
	appoint3 := `https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxd09c7682905819e6&redirect_uri=http%3a%2f%2fwww.elepick.com%2fapi%2fbranch&response_type=code&scope=snsapi_base&state=123#wechat_redirect`
	bt1, bt2, bt3 := menu.NewViewButton("预约体检", appoint), menu.NewViewButton("我的报告", "http://www.elepick.com/api"), menu.NewViewButton("分院", appoint3)
	button := menu.NewTopButton("迪安").AddSub(bt1).AddSub(bt2).AddSub(bt3)
	return menu.CreateMenu(button, access_token)
}
