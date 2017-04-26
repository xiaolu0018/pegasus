package handler

import (
	"net/http"

	"192.168.199.199/bjdaos/pegasus/pkg/appoint/appointment"
	"github.com/julienschmidt/httprouter"
)

func CreateHttpRouter() http.Handler {
	r := httprouter.New()

	r.GET("/api/organazations", ListHandler)
	r.PUT("/api/organazation/:code/config/basic", CreateBasicHandler)
	r.POST("/api/organazation/:code/config/special", CreateSpecialHandler)

	r.POST("/api/appointment", appointment.CreateAppointmentHandler)
	r.POST("/api/appointment/:appointid/cancel", appointment.CancelAppointmentHandler)
	r.POST("/api/appointment/:appointid/comment", appointment.CreateCommentHandler)
	r.PUT("/api/appointment", appointment.UpdateAppointmentHandler)
	r.GET("/api/appointment/:appointid", appointment.GetAppointmentHandler)
	r.GET("/api/appointmenlist", appointment.ListAppointmentsHandler)

	r.GET("/api/pinto/checkups", ListCheckupHandler)
	return r
}
