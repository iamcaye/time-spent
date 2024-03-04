package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iamcaye/time-spent/api"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api.SetupRouter(r)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
