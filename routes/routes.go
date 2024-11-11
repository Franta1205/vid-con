package routes

import (
	"github.com/gin-gonic/gin"
	"vid-con/controllers"
)

func SetRoutes(r *gin.Engine) {
	homeController := controllers.InitHone()
	r.GET("/", homeController.Home)

	videoController := controllers.InitVideo()
	callGroup := r.Group("/call")
	{
		callGroup.GET("/:room_id", videoController.Show)
		callGroup.POST("/start", videoController.Start)
		callGroup.POST("/join", videoController.Join)
	}
}
