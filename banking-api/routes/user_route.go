package routes

import (
	"banking-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	//All routes related to users comes here
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:id", controllers.GetAUser())
	router.PUT("/user/:id", controllers.EditUserBalance())
	router.DELETE("/user/:id", controllers.DeleteAUser())
	router.GET("/users", controllers.GetAllUsers())
}
