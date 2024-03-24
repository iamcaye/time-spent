package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iamcaye/time-spent/internal/events"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	{
		eventRouter := events.NewEventRouter()
		eventsGroup := api.Group("/events")
		{
			eventRouter.SetupRouter(eventsGroup)
			fmt.Println("Event routes created")
		}
	}

}
