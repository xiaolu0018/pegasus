package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateHttpRouter() http.Handler {
	r := httprouter.New()

	r.POST("/api/appointment/person", BookPersonHandler)
	r.POST("/api/appointment/plan", BookPlanHandler)
	r.GET("/api/examination/status", GetExamStatusHandler)
	r.POST("/api/examination/cancel", CancelExamHandler)
	r.GET("/api/sales/checkups", GetCheckupCodesBySaleCodesHandler)

	return r
}
