package api_token

import (
	"sync"
	"time"

	"github.com/1851616111/util/http"
	"github.com/golang/glog"
)

var DefaultGrantType string = "client_credential"
var TokenUrl string = "https://api.weixin.qq.com/cgi-bin/token"
var TicketUrl string = "https://api.weixin.qq.com/cgi-bin/ticket/getticket"
var TokenCtrl *Controller
var once sync.Once

//curl -G "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxd09c7682905819e6&secret=b9938ddfec045280eba89fab597a0c41"
func InitController(appID string, secret string) *Controller {
	once.Do(func() {
		TokenCtrl = &Controller{
			l:      sync.RWMutex{},
			params: http.NewParams().Add("grant_type", "client_credential").Add("appid", appID).Add("secret", secret),
		}
	})

	return TokenCtrl
}

func (c *Controller) Run() error {
	if err := c.updateToken(); err != nil {
		return err
	}
	glog.Infof("pkg.access_token.updateToken: first update token(%s) success", c.token)

	if err := c.updateTicket(); err != nil {
		return err
	}
	glog.Infof("pkg.access_token.updateTicket: first update ticket(%s) success", c.ticket)

	go func() {
		for {
			//提前60秒进行更新
			time.Sleep(time.Second * time.Duration(c.getExpire()))
			//time.Sleep(time.Second*time.Duration(20 + 60))

			if err := c.updateToken(); err != nil {
				glog.Errorf("pkg.access_token.updateToken: continue update token err %v\n", err)

				c.setExpire(20 + 60) //actual 20 second
				c.err = err
			} else {
				glog.Infof("pkg.access_token.updateToken: update token(%s) success", c.token)
			}

			if err := c.updateTicket(); err != nil {
				glog.Errorf("pkg.access_token.updateTicket: continue ticket token err %v\n", err)
			} else {
				glog.Infof("pkg.access_token.updateTicket: update ticket(%s) success", c.ticket)
			}
		}
	}()

	return nil
}
