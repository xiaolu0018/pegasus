package api_token

import (
	"sync"
	"time"

	"github.com/1851616111/util/http"
	"github.com/golang/glog"
)

var DefaultGrantType string = "client_credential"
var TokenUrl string = "https://api.weixin.qq.com/cgi-bin/token"

//curl -G "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxd09c7682905819e6&secret=b9938ddfec045280eba89fab597a0c41"
func NewController(appID string, secret string) *Controller {
	return &Controller{
		l:      sync.RWMutex{},
		params: http.NewParams().Add("grant_type", "client_credential").Add("appid", appID).Add("secret", secret),
	}
}

func (c *Controller) Run() error {
	if err := c.updateToken(); err != nil {
		return err
	}

	glog.Infof("pkg.access_token.updateToken: first update token(%s) success", c.token)
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
		}
	}()

	return nil
}
