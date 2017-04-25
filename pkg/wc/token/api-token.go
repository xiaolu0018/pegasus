package token

import (
	wctoken "github.com/1851616111/util/weichat/util/api-token"
	wchandler "github.com/1851616111/util/weichat/util/handler"
	usertoken "github.com/1851616111/util/weichat/util/user-token"
)

var TokenCtrl *wctoken.Controller

func InitApiToken(appID, appSecret string) error {
	wchandler.APP_ID = appID
	wchandler.Token = usertoken.NewTokenConfig(appID, appSecret)

	return startTokenCtrl(appID, appSecret)
}

func startTokenCtrl(appID, appSecret string) error {
	TokenCtrl = wctoken.NewController(appID, appSecret)
	return TokenCtrl.Run()
}
