package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadEventsRoutes(r *gin.RouterGroup) {
	events := r.Group("/events")
	{
		events.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

}
