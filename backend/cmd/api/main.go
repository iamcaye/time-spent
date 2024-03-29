package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iamcaye/time-spent/internal/routes"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func main() {
	fmt.Println("Starting the application...")
	start_time := time.Now()
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routes.SetupRouter(r)

	err = r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		fmt.Println("An error occurred while setting trusted proxies")
	}

	port := os.Getenv("PORT")
	elapsed := time.Since(start_time)
	fmt.Println("[TIME] Start the application: ", elapsed)
	error := r.Run(fmt.Sprintf(":%s", port))
	if error != nil {
		fmt.Println("An error occurred while starting the application")
	}

}
