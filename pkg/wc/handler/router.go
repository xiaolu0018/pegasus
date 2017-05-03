package handler

import (
	"net/http"
	"path/filepath"

	"192.168.199.199/bjdaos/pegasus/pkg/wc/appointment"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/banner"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/branch"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/capacitymanage"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/image"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/plan"
	"192.168.199.199/bjdaos/pegasus/pkg/wc/user"
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

	r.PUT("/api/banner", authAdmin(banner.UpsertHandler))
	r.GET("/api/banners", banner.GetHandler)

	r.PUT("/api/plan", authAdmin(plan.UpsertHandler))
	r.GET("/api/plans", plan.GetPlansHandler)

	r.PUT("/api/appointment", authUser(appointment.CreateHandler))
	//r.POST("/api/appointment/:id/branch", authUser(appointment.UpdateHandler))
	r.POST("/api/appointment/:appointid/confirm", authUser(appointment.ConfirmCreatHandler))
	r.GET("/api/appointment/:appointid/confirm", authUser(appointment.GetAppointmentConfirmHandler))
	r.POST("/api/appointments/:id/cancel", authUser(appointment.CancelHandler))
	r.GET("/api/appointments", authUser(appointment.ListAppointmentHandler))
	r.POST("/api/appointment/:appointid/comment", authUser(appointment.CreateCommentHandler))
	r.POST("/api/report/:mobile", authUser(appointment.GetCheckNoForReport))

	r.GET("/api/branch/:id/offday", capacitymanage.GetOffDaysHandler)
	r.POST("/api/manage/branch", authAdmin(branch.CreateHandler))
	r.PUT("/api/manage/branch/:id", authAdmin(branch.UpdateHandler))
	r.GET("/api/branches", branch.ListHandler)
	r.POST("/api/manage/uploadfile", authAdmin(image.SaveImageHandler))

	//r.GET("/api/appoin/test", appointment.ConfirmAppointment)

	dist, _ = filepath.Abs(dist)
	r.ServeFiles("/dist/*filepath", http.Dir(dist))

	return r
}
