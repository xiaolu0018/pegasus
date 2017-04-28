package manager

import (
	"github.com/1851616111/util/weichat/event"
	"github.com/1851616111/util/weichat/handler"
)

func CreateActivity() error {
	article := event.NewArticleAction("第一次订阅推送测试", "第一次订阅推送测试", "http://wx.qlogo.cn/mmopen/PiajxSqBRaELA3P616lZg1b33Vzv8Hhc6GnPEOmUZG7Gl8sL4aNic6WhOmezogpWmoXad53BuoVXmPHTsnVLbbvA/0", "http://www.elepick.com:8080/swagger-ui/dist/")
	handler.EventManager.Registe(event.E_Subscribe, article)
	handler.EventManager.Registe(event.E_UnSubscribe, article)
	return nil
}
