package api_token

import (
	"github.com/1851616111/util/http"
	"sync"
)

type Controller struct {
	l sync.RWMutex //protect token and config

	token     string
	expireSec uint16
	params    *http.Params //appid,secret and grant_type

	ticket          string
	ticketExpireSec uint16
	err             error //unknown error
}

type token struct {
	Token  string `json:"access_token,omitempty"`
	Expire uint16 `json:"expires_in,omitempty"`
	Code   int    `json:"errcode,omitempty"`
	Msg    string `json:"errmsg,omitempty"`
}

//{"errcode":0,"errmsg":"ok","ticket":"kgt8ON7yVITDhtdwci0qeYgh_SptvWn_34kRNSvKy3RKs3wCJoxZ5zcDzXF7Hw9NMTaIvk3V0PYbxVNrJT-Rmw","expires_in":7200}
type ticketMsg struct {
	Code   int    `json:"errcode"`
	Msg    string `json:"errmsg"`
	Ticket string `json:"ticket"`
	Expire int    `json:"expires_in"`
}
