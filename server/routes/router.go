package routes

import (
	"net/http"
	"time"

	"github.com/1pedrohfreitas/pcams_back_go/controllers"
	"github.com/1pedrohfreitas/pcams_back_go/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"Origin", "*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
	}))
	router.StaticFS("api/v1/static", http.Dir("static"))
	main := router.Group("api/v1")
	{

		login := main.Group("login")
		{
			login.POST("/", controllers.Login)
			login.POST("/validatetoken", controllers.ValidateToken)
		}
		users := main.Group("users", middlewares.Auth())
		{
			users.GET("/:id", controllers.ShowUser)
			users.POST("/", controllers.CreateUser)
			users.GET("/", controllers.ShowUsers)
			users.PUT("/", controllers.UpdateUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}
		cams := main.Group("cams", middlewares.Auth())
		{
			cams.GET("/:id", controllers.ShowCam)
			cams.POST("/", controllers.CreateCam)
			cams.GET("/", controllers.ShowCams)
			cams.PUT("/", controllers.UpdateCam)
			cams.DELETE("/:id", controllers.DeleteCam)
			cams.GET("/mycams/:userid", controllers.ShowCamsByUser)
			cams.POST("/camaccesspermission", controllers.CreateCamAccessPermission)
		}
	}
	return router
}
