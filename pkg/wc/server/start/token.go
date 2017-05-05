package start

import (
	wchandler "github.com/1851616111/util/weichat/handler"
	apitoken "github.com/1851616111/util/weichat/util/api-token"
	usertoken "github.com/1851616111/util/weichat/util/user-token"
)

func (o *ActivityConfig) initApiToken() error {
	wchandler.APP_Sign_Token = o.AppSignToken
	wchandler.Token = usertoken.NewTokenConfig(o.AppID, o.AppSecret)

	apitoken.InitController(o.AppID, o.AppSecret)
	return apitoken.TokenCtrl.Run()
}
