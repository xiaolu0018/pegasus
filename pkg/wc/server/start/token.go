package start

import (
	wchandler "github.com/1851616111/util/weichat/handler"
	wcpitoken "github.com/1851616111/util/weichat/util/api-token"
	usertoken "github.com/1851616111/util/weichat/util/user-token"
)

func (o *ActivityConfig) initApiToken() error {
	wchandler.APP_ID = o.AppID
	wchandler.Token = usertoken.NewTokenConfig(o.AppID, o.AppSecret)

	wcpitoken.InitController(o.AppID, o.AppSecret)
	return wcpitoken.TokenCtrl.Run()
}
