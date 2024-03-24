package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/iamcaye/time-spent/internal/routes"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	fmt.Println("Starting the application...")
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routes.SetupRouter(r)

	r.SetTrustedProxies([]string{"127.0.0.1"})

	port := os.Getenv("PORT")
	error := r.Run(fmt.Sprintf(":%s", port))
	if error != nil {
		fmt.Println("An error occurred while starting the application")
	}

}
