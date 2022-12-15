package main

import (
	routes "jwt_with_gin/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT") //env variable fetch

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger()) //beautiful logger
	//2 type of route

	routes.AuthRoutes(router) //passing router to routes
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"Success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"Success": "Access granted for api-2"})
	})

	router.Run(":" + port) //similar to listenAndServe
}
