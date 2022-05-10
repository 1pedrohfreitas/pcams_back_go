package routes

import (
	"github.com/1pedrohfreitas/pcams_back_go/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.GET("/:id", controllers.ShowUser)
			users.POST("/", controllers.CreateUser)
			users.GET("/", controllers.ShowUsers)
			users.PUT("/", controllers.UpdateUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}
		cams := main.Group("cams")
		{
			cams.GET("/:id", controllers.ShowUser)
			cams.POST("/", controllers.CreateUser)
			cams.GET("/", controllers.ShowUsers)
			cams.PUT("/", controllers.UpdateUser)
			cams.DELETE("/:id", controllers.DeleteUser)
		}
	}
	return router
}
