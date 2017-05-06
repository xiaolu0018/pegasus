package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateHttpRouter() http.Handler {
	r := httprouter.New()

	r.POST("/api/create/bookrecord", CreateBookRecordHandler)
	r.POST("/api/create/examwithplan", CreateExamsHandler)
	r.GET("/api/examination/status", GetExamStatusHandler)
	r.POST("/api/cancel/exam",CancelExamHandler)
	r.GET("/api/sales/checkups", GetCheckupCodesBySaleCodesHandler)

	return r
}
