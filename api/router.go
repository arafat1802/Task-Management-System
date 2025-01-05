package routes

import (
	"github.com/arafat1802/Task-Management-System/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users", controllers.GetUsers)
	r.POST("/createUser", controllers.CreateUser)

	return r
}
