package manager

import (
	"github.com/1851616111/util/weichat/event"
	"github.com/1851616111/util/weichat/handler"
	"fmt"
)

func WatchEvent(appid, schema, domain string) error {
	url := fmt.Sprintf(`https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s%%3a%%2f%%2f%s%%2fapi%%2factivity%%2findex&response_type=code&scope=snsapi_userinfo&state=123#wechat_redirect`,
		appid, schema, domain)
	imgUrl := fmt.Sprintf("%s://%s/dist/activity/img/head.png")
	article := event.NewArticleAction("跟衰老 Say Bye Bye", "晒合影, 喊好友来助力！！！",imgUrl, url)
	handler.EventManager.Registe(event.E_Subscribe, article)
	handler.EventManager.Registe(event.E_UnSubscribe, article)
	return nil
}
