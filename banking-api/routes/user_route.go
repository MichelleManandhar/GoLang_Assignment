package routes

import (
	"banking-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	//All routes related to users comes here
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:ids", controllers.GetAUser())
	router.PUT("/user/:ids", controllers.EditUserBalance())
	router.DELETE("/user/:ids", controllers.DeleteAUser())
	router.GET("/users", controllers.GetAllUsers())
	router.PUT("user/withdraw/:ids", controllers.WithdrawBalance())
	router.PUT("user/deposit/:ids", controllers.DepositBalance())
}
