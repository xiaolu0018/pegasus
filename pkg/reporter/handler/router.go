package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateHttpRouter() http.Handler {
	r := httprouter.New()
	r.GET("/api/report", GetReport)
	r.GET("/api/report/list", AuthHandler(ReportListHandler))
	r.POST("/api/report/status", AuthHandler(UpdateStatusHandler))
	return r
}
