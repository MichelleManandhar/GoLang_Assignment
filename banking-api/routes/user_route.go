package routes

import (
	"banking-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	//All routes related to users comes here
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:acc", controllers.GetAUser())
	router.PUT("/user/:acc", controllers.EditUserBalance())
	router.DELETE("/user/:acc", controllers.DeleteAUser())
	router.GET("/users", controllers.GetAllUsers())
	router.PUT("user/withdraw/:acc", controllers.WithdrawBalance())
	router.PUT("user/deposit/:acc", controllers.DepositBalance())
}
