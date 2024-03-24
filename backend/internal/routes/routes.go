package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iamcaye/time-spent/db"
	"github.com/iamcaye/time-spent/internal/events"
)

func SetupRouter(r *gin.Engine) {
	DB := db.InitDB()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	{
		eventRouter := events.NewEventRouter(DB)
		eventsGroup := api.Group("/events")
		{
			eventRouter.SetupRouter(eventsGroup)
		}
	}

}
