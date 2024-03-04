package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iamcaye/time-spent/api/events"
)

func SetupRouter(r *gin.Engine) {
	api := r.Group("/api")
	events.LoadEventsRoutes(api)
}
