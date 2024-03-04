package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/dejauls/task-5-pbi-fullstack-developer--Rajaul-Bani-Safar-/controllers"
	"github.com/dejauls/task-5-pbi-fullstack-developer--Rajaul-Bani-Safar-/middlewares"
)

func SetupRoutes(router *gin.Engine) {
	publicGroup := router.Group("/")
	{
		publicGroup.POST("/users/register", controllers.RegisterUser)
		publicGroup.POST("/users/login", controllers.LoginUser)
	}

	userAuthGroup := router.Group("/users").Use(middlewares.AuthMiddleware())
	{
		userAuthGroup.PUT("/:userId", controllers.UpdateUser)
		userAuthGroup.DELETE("/:userId", controllers.DeleteUser)
	}

	photoAuthGroup := router.Group("/photos").Use(middlewares.AuthMiddleware())
	{
		photoAuthGroup.POST("/", controllers.CreatePhoto)
		photoAuthGroup.GET("/", controllers.GetPhotos)
		photoAuthGroup.PUT("/:photoId", controllers.UpdatePhoto)
		photoAuthGroup.DELETE("/:photoId", controllers.DeletePhoto)
	}
}
	