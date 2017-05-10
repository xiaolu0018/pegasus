package manager

import (
	"github.com/1851616111/util/weichat/event"
	"github.com/1851616111/util/weichat/handler"
)

func WatchEvent(appid, schema, domain string) error {
	url := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxadb8865cdb995ded&redirect_uri=http://hd1.dahe100.cn/api/activity/index&response_type=code&scope=snsapi_userinfo&state=123#wechat_redirect"

	imgUrl := "http://hd1.dahe100.cn/dist/activity/img/head.png"
	article := event.NewArticleAction("跟衰老 Say Bye Bye", "晒合影, 喊好友来助力！！！", imgUrl, url)
	handler.EventManager.Registe(event.E_Subscribe, article)
	handler.EventManager.Registe(event.E_UnSubscribe, article)
	return nil
}
