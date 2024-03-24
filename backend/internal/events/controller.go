package events

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type EventController struct {
	service *EventService
}

func (ctl *EventController) GetEvents(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": ctl.service.GetEvents(),
	})
}

func (ctl *EventController) GetEventByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	c.JSON(200, gin.H{
		"data": ctl.service.GetEventByID(id),
	})
}

func (ctl *EventController) CreateEvent(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": ctl.service.CreateEvent(Event{}),
	})
}

func (ctl *EventController) UpdateEvent(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": ctl.service.UpdateEvent(Event{}),
	})
}

func (ctl *EventController) DeleteEvent(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	c.JSON(200, gin.H{
		"data": ctl.service.DeleteEvent(id),
	})
}
