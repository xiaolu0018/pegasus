//==================================================================
//创建时间：2017-4-23 首次创建
//功能描述：router
//创建人：郭世江
//修改记录：若要修改请记录
//==================================================================

package handler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func CreateHttpRouter() http.Handler {
	r := httprouter.New()
	r.GET("/api/heart/login", httprouter.Handle(GetAccountInfo))
	r.GET("/api/heart/queryCode", httprouter.Handle(GetQueryCode))
	r.GET("/api/heart/data", httprouter.Handle(GetHeartData))
	r.GET("/api/heart/updateResult", httprouter.Handle(UpdateResult))
	//r.GET("/api/heart/queue", httprouter.Handle(GetNextRoom))
	//r.GET("api/heart/checkupItem", httprouter.Handle(GetPersonCheckUpItem))
	return r
}
