package routes

import (
	"github.com/gin-gonic/gin"
	"vid-con/controllers"
)

func SetRoutes(r *gin.Engine) {
	homeController := controllers.InitHone()
	r.GET("/", homeController.Index)

	videoController := controllers.InitVideo()
	callGroup := r.Group("/call")
	{
		callGroup.GET("/:room_id", videoController.Show)
		callGroup.POST("/start", videoController.Start)
		callGroup.GET("/confirm_start", videoController.ConfirmStart)
		callGroup.POST("/join", videoController.Join)
		callGroup.GET("/confirm_join", videoController.ConfirmJoin)
		callGroup.POST("/leave", videoController.Leave)
	}
}
