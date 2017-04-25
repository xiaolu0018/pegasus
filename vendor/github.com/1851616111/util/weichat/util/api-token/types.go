package api_token

import (
	"github.com/1851616111/util/http"
	"sync"
)

type Controller struct {
	l sync.RWMutex //protect token and config

	expireSec uint16
	params    *http.Params //appid,secret and grant_type

	token string
	err   error //unknown error
}

type token struct {
	Token  string `json:"access_token,omitempty"`
	Expire uint16 `json:"expires_in,omitempty"`
	Code   int    `json:"errcode,omitempty"`
	Msg    string `json:"errmsg,omitempty"`
}
