package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateHttpRouter() http.Handler {
	r := httprouter.New()
	r.GET("/api/report", GetReport)
	r.GET("/api/report/list", AuthHandler(ReportListHandler))

	//这里应该是POST, 但是客户端发post不太会
	r.GET("/api/report/status", AuthHandler(UpdateStatusHandler))
	return r
}
