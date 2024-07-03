package routes

import (
	"pbi-task/controllers"
	"pbi-task/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/api/user/register", controllers.RegisterUser, controllers.GenerateToken)
	router.POST("/api/user/login", controllers.UserLogin, controllers.GenerateToken)
	router.Use(middleware.Auth()).GET("/ping", controllers.Ping)

	api := router.Group("/api/user")
	{
		api.GET("/", controllers.GetUser, middleware.AuthUser)
		api.GET("/id", controllers.GetUserById, middleware.AuthUser)
		api.PUT("/id", controllers.UserUpdate, middleware.AuthUser)
		api.DELETE("/id", controllers.UserDelete)

	}
	return router
}
