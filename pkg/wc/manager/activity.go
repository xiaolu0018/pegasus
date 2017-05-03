package manager

import (
	"github.com/1851616111/util/weichat/event"
	"github.com/1851616111/util/weichat/handler"
)

func WatchEvent() error {
	article := event.NewArticleAction("跟衰老 Say Bye Bye", "晒合影, 喊好友来助力！！！", "http://www.elepick.com/dist/activity/img/head.png", "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxd09c7682905819e6&redirect_uri=http%3a%2f%2fwww.elepick.com%2fapi%2factivity%2findex&response_type=code&scope=snsapi_userinfo&state=123#wechat_redirect")
	handler.EventManager.Registe(event.E_Subscribe, article)
	handler.EventManager.Registe(event.E_UnSubscribe, article)
	return nil
}
