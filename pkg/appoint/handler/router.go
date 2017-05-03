package handler

import (
	"net/http"

	"192.168.199.199/bjdaos/pegasus/pkg/appoint/appointment"
	"192.168.199.199/bjdaos/pegasus/pkg/appoint/login"
	"github.com/julienschmidt/httprouter"
)

func CreateHttpRouter() http.Handler {
	r := httprouter.New()
	//r.GET("/api/index",GetIndexHandler)

	r.POST("/api/login", login.CheckLoginHandler)
	r.POST("/api/login/changepwd", login.ChangeLoginHandler)

	r.GET("/api/organazations", ListHandler)
	r.GET("/api/organizations/wc", ListHandlerWC)
	r.POST("/api/organazation/:code/config/basic", CreateBasicHandler)
	r.GET("/api/organazation/:code", GetBasicHandler)
	r.POST("/api/organazation/:code/config/special", CreateSpecialHandler)

	r.GET("/api/plans", GetPlansHandler)
	r.GET("/api/banners", GetBannersHandler)
	r.GET("/api/offday/:code", GetOffDayHandler)

	r.POST("/api/appointment", appointment.CreateAppointmentHandler)
	r.POST("/api/appointment/:appointid/cancel", appointment.CancelAppointmentHandler)
	r.POST("/api/appointment/:appointid/comment", appointment.CreateCommentHandler)
	r.PUT("/api/appointment", appointment.UpdateAppointmentHandler)
	r.GET("/api/appointment/:appointid", appointment.GetAppointmentHandler)
	r.GET("/api/appointmentlist", appointment.ListAppointmentsHandler)
	r.GET("/api/appointmentlist/wc", appointment.ListAppointmentsForWeChatHandler)
	r.GET("/api/pinto/checkups", ListCheckupHandler)
	return r
}
