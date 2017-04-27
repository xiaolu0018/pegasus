package handler

import (
	"github.com/1851616111/util/rand"
	tk "github.com/1851616111/util/weichat/util/user-token"
	"github.com/julienschmidt/httprouter"

	"192.168.199.199/bjdaos/pegasus/pkg/wc/user"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/util"
)

const TO_CACHE_SCOPE = "snsapi_userinfo"

func completeUserInfo(t *tk.Token, ps *httprouter.Params) error {
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
