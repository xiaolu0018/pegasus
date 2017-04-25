package user_token

import (
	"encoding/json"
	"github.com/1851616111/util/http"
)

//details https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140842&token=&lang=zh_CN
var AccessTokenAddress string = "https://api.weixin.qq.com/sns/oauth2/access_token"

type Config struct {
	Address   string
	AppID     string
	Secret    string
	GrantType string
}

func NewTokenConfig(app_id, secret string) *Config {
	return &Config{
		Address:   AccessTokenAddress,
		AppID:     app_id,
		Secret:    secret,
		GrantType: "authorization_code",
	}
}

func (c *Config) Exchange(code string) (*Token, error) {
	rsp, err := http.Send(&http.HttpSpec{
		URL:    c.Address,
		Method: "GET",
		URLParams: http.NewParams().Add("appid", c.AppID).Add("secret", c.Secret).
			Add("code", code).Add("grant_type", c.GrantType),
	})

	if err != nil {
		return nil, err
	}

	var token Token
	if err = json.NewDecoder(rsp.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}

type Token struct {
	Access_Token  string `json:"access_token"`
	Expire_In     int    `json:"expires_in"`
	Refresh_Token string `json:"refresh_token"`
	Open_ID       string `json:"openid"`
	Scope         string `json:"scope"`
}
