package manager

import (
	"fmt"

	"github.com/1851616111/util/weichat/manager/menu"
	wcpitoken "github.com/1851616111/util/weichat/util/api-token"
)

func CreateMenuForActivity(appid, schema, domain string) error {
	appoint2 := fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s%%3a%%2f%%2f%s%%2fapi%%2factivity%%2findex&response_type=code&scope=snsapi_userinfo&state=123#wechat_redirect",
		appid, schema, domain)

	bt1 := menu.NewViewButton("母亲节活动", appoint2)
	//appoint := `https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxd09c7682905819e6&redirect_uri=http%3a%2f%2fwww.elepick.com%2fapi%2fappoint&response_type=code&scope=snsapi_userinfo&state=123#wechat_redirect`
	//appoint3 := `https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxd09c7682905819e6&redirect_uri=http%3a%2f%2fwww.elepick.com%2fapi%2fbranch&response_type=code&scope=snsapi_userinfo&state=123#wechat_redirect`
	//appoint2 := `https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxd09c7682905819e6&redirect_uri=http%3a%2f%2fwww.elepick.com%2fapi%2freportmenu&response_type=code&scope=snsapi_userinfo&state=123#wechat_redirect`
	//bt1, bt2, bt3 := menu.NewViewButton("预约体检", appoint), menu.NewViewButton("我的报告", appoint2), menu.NewViewButton("分院", appoint3)
	//button := menu.NewTopButton("迪安").AddSub(bt1).AddSub(bt2).AddSub(bt3)

	button := menu.NewTopButton("活动").AddSub(bt1)
	return menu.CreateMenu(button, wcpitoken.TokenCtrl.GetToken())
}

func CreateMenu() error {
	appoint := `https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxd09c7682905819e6&redirect_uri=http%3a%2f%2fwww.elepick.com%2fapi%2fappoint&response_type=code&scope=snsapi_userinfo&state=123#wechat_redirect`
	appoint3 := `https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxd09c7682905819e6&redirect_uri=http%3a%2f%2fwww.elepick.com%2fapi%2fbranch&response_type=code&scope=snsapi_userinfo&state=123#wechat_redirect`
	appoint2 := `https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxd09c7682905819e6&redirect_uri=http%3a%2f%2fwww.elepick.com%2fapi%2freportmenu&response_type=code&scope=snsapi_userinfo&state=123#wechat_redirect`
	bt1, bt2, bt3 := menu.NewViewButton("预约体检", appoint), menu.NewViewButton("我的报告", appoint2), menu.NewViewButton("分院", appoint3)
	button := menu.NewTopButton("迪安").AddSub(bt1).AddSub(bt2).AddSub(bt3)
	return menu.CreateMenu(button, wcpitoken.TokenCtrl.GetToken())
}
