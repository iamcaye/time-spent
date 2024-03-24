package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iamcaye/time-spent/internal/routes"
)

func main() {
	fmt.Println("Starting the application...")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routes.SetupRouter(r)

	error := r.Run(":8080")
	if error != nil {
		fmt.Println("An error occurred while starting the application")
	}

}
