package router

import (

	"github.com/gin-gonic/gin"
	"github.com/levifleal/MyFlix/handler"
)

func initRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/serie", handler.ShowSerieHandler)
		v1.POST("/serie", handler.CreateSerieHandler)
		v1.DELETE("/serie", handler.DeleteSerieHandler)
		v1.PUT("/serie", handler.UpdateSerieHandler)
		v1.GET("/series", handler.ListSeriesHandler)
	}
}
