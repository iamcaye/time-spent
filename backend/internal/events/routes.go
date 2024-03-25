package events

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type EventRouter struct {
	ctl *EventController
}

func NewEventRouter(db *sqlx.DB) *EventRouter {
	return &EventRouter{
		ctl: &EventController{
			service: &EventService{
				db: db,
			},
		},
	}
}

func (router *EventRouter) SetupRouter(r *gin.RouterGroup) {
	r.GET("", router.ctl.GetEvents)
	r.GET(":id", router.ctl.GetEventByID)
	r.POST("", router.ctl.CreateEvent)
	r.PUT(":id", router.ctl.UpdateEvent)
	r.DELETE(":id", router.ctl.DeleteEvent)

	r.GET(":id/start", router.ctl.StartEvent)
	r.GET(":id/stop", router.ctl.StopEvent)
}
