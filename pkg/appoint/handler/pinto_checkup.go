package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"192.168.199.199/bjdaos/pegasus/pkg/appoint/cache"
	"192.168.199.199/bjdaos/pegasus/pkg/appoint/pinto"
	httputil "github.com/1851616111/util/http"
)

func ListCheckupHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ck, err := cache.Get("inner_system", pinto.CACHE_CHECKUP)
	if err != nil {
		httputil.Response(w, 400, err)
		return
	}

	httputil.ResponseJson(w, 200, ck)
}