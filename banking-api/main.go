package main

import (
	"banking-api/configs"
	"banking-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//test in postman
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"data": "Hello from Gin-gonic & mongoDB",
	// 	})
	// })

	//Run database
	configs.ConnectDB()

	//Routes
	routes.UserRoute(router)

	router.Run("localhost:8080")
}
