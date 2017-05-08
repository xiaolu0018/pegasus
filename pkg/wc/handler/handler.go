package handler

import (
	"bjdaos/pegasus/pkg/wc/user"
	"bjdaos/pegasus/pkg/wc/util"
	"github.com/1851616111/util/rand"
	"github.com/julienschmidt/httprouter"

	tk "github.com/1851616111/util/weichat/util/user-token"
)

const TO_CACHE_SCOPE = "snsapi_userinfo"

func CompleteAccessTokenInfo(t *tk.Token, ps *httprouter.Params) error {
	var token string
	var ok bool
	var id string = t.Open_ID

	if ok, token = user.IDCache.GetWorkingToken(id); !ok {

		token = rand.String(user.TokenLength)
		user.IDCache.CacheSysToken(id, token)
	}
	if t.Scope == TO_CACHE_SCOPE {
		user.IDCache.CacheWCToken(id, t.Access_Token)
	}

	newPs := util.AddParam(*ps, "bear_token", token)
	*ps = newPs

	return nil
}

func CompleteOpenidInfo(t *tk.Token, ps *httprouter.Params) error {
	newPs := util.AddParam(*ps, "openid", t.Open_ID)
	*ps = newPs

	return nil
}
