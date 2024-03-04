package events

import (
	"github.com/gin-gonic/gin"
)

func LoadEventsRoutes(r *gin.RouterGroup) {
	r.GET("/events", GetAllEventsCtl)
	r.GET("/events/:id", GetEventCtl)
	r.POST("/events", SaveEventCtl)
}
