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
	event := Event{}
	err := c.BindJSON(&event)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": ctl.service.CreateEvent(event),
	})
}

func (ctl *EventController) UpdateEvent(c *gin.Context) {
	event := Event{}
	err := c.BindJSON(&event)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": ctl.service.UpdateEvent(event),
	})
}

func (ctl *EventController) DeleteEvent(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	c.JSON(200, gin.H{
		"data": ctl.service.DeleteEvent(id),
	})
}

func (ctl *EventController) StartEvent(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	c.JSON(200, gin.H{
		"data": ctl.service.StartEvent(id),
	})
}

func (ctl *EventController) StopEvent(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	c.JSON(200, gin.H{
		"data": ctl.service.StopEvent(id),
	})
}
