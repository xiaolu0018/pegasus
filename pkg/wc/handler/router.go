package handler

import (
	"net/http"
	"path/filepath"

	"bjdaos/pegasus/pkg/wc/appointment"
	"bjdaos/pegasus/pkg/wc/banner"
	"bjdaos/pegasus/pkg/wc/branch"
	"bjdaos/pegasus/pkg/wc/capacitymanage"
	"bjdaos/pegasus/pkg/wc/plan"
	"bjdaos/pegasus/pkg/wc/user"
	"github.com/1851616111/util/weichat/handler"
	"github.com/julienschmidt/httprouter"
)

func AddApiToRouter(r *httprouter.Router, dist string) http.Handler {
	r.GET("/api/basic/signature", handler.DeveloperValidater)
	r.POST("/api/basic/signature", handler.EventAction)

	r.POST("/api/user", authUser(user.UpsertInfoHandler))
	r.POST("/api/user/label", authUser(user.UpdateLabelHandler))
	r.GET("/api/user/label", authUser(user.GetLabelHandler))
	r.GET("/api/user", authUser(user.GetHandler))

	r.GET("/api/banners", banner.GetHandler)
	r.GET("/api/plans", plan.GetPlansHandler)

	r.PUT("/api/appointment", authUser(appointment.CreateHandler))
	r.POST("/api/appointment/:appointid/confirm", authUser(appointment.ConfirmCreatHandler))
	r.GET("/api/appointment/:appointid/confirm", authUser(appointment.GetAppointmentConfirmHandler))
	r.POST("/api/appointments/:id/cancel", authUser(appointment.CancelHandler))
	r.GET("/api/appointments", authUser(appointment.ListAppointmentHandler))
	r.POST("/api/appointment/:appointid/comment", authUser(appointment.CreateCommentHandler))
	r.POST("/api/report/:mobile", authUser(appointment.GetCheckNoForReport))
	r.GET("/api/appoint/report/:checkno/:appid", authUser(appointment.GetReportByAppid))

	r.GET("/api/branch/:id/offday", capacitymanage.GetOffDaysHandler)
	r.GET("/api/branches", branch.ListHandler)

	dist, _ = filepath.Abs(dist)
	r.ServeFiles("/dist/*filepath", http.Dir(dist))

	return r
}
