package events

import (
	"github.com/gin-gonic/gin"
)

type EventRouter struct {
	ctl *EventController
}

func NewEventRouter() *EventRouter {
	return &EventRouter{
		ctl: &EventController{
			service: &EventService{},
		},
	}
}

func (router *EventRouter) SetupRouter(r *gin.RouterGroup) {
	r.GET("", router.ctl.GetEvents)
	r.GET(":id", router.ctl.GetEventByID)
	r.POST("", router.ctl.CreateEvent)
	r.PUT("", router.ctl.UpdateEvent)
	r.DELETE(":id", router.ctl.DeleteEvent)
}
