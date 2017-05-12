package manager

import (

	"github.com/1851616111/util/weichat/manager/menu"
	wcpitoken "github.com/1851616111/util/weichat/util/api-token"
	"fmt"
)

func CreateMenu(appid, schema, domain string) error {
	bt_1_1 := menu.NewViewButton("免费wifi", "http://wifi.weixin.qq.com/mbl/connect.xhtml?type=1")
	bt_1_2 := menu.NewViewButton("查看报告", "http://hd1.dahe100.cn/dist/activity/updating.html")
	bt_1_3 := menu.NewViewButton("预约", "http://hd1.dahe100.cn/dist/activity/updating.html")
	bt_1 := menu.NewTopButton("体检服务").AddSub(bt_1_3).AddSub(bt_1_2).AddSub(bt_1_1)

	appoint2 := fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s%%3a%%2f%%2f%s%%2fapi%%2factivity%%2findex&response_type=code&scope=snsapi_userinfo&state=123#wechat_redirect",
		appid, schema, domain)
	bt_2_1 := menu.NewViewButton("母亲节活动", appoint2)
	bt_2 := menu.NewTopButton("活动").AddSub(bt_2_1)

	bt_3 := menu.NewTopButton("关于我们")
	bt_3_2 := menu.NewViewButton("公司简介", "http://mp.weixin.qq.com/s?__biz=MzAwNDE4OTgyNw==&mid=506735864&idx=1&sn=ccb723be42baa87a748a4f511a2f8260&chksm=00ef37143798be02f90644c35eb2a0e95e4b38b6c181c0711506d0ca8874039328a0b46c9fbf&scene=18#wechat_redirect")
	Bt_3_1 := menu.NewViewButton("健检动态", "http://mp.weixin.qq.com/s?__biz=MzAwNDE4OTgyNw==&mid=506735864&idx=2&sn=1fb3905e4d2d9961d4a00832953e5284&chksm=00ef37143798be02c47519149246e10a8cf0f4eb5c8d892f51d3a74312235c4d7bc8ff929d3f&scene=18#wechat_redirect")
	bt_3.AddSub(bt_3_2).AddSub(Bt_3_1)

	return menu.CreateMenu(wcpitoken.TokenCtrl.GetToken(), bt_1, bt_2, bt_3)
	//return menu.CreateMenu(wcpitoken.TokenCtrl.GetToken(), bt_1, bt_3)
}
