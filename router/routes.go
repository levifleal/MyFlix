package router

import (

	"github.com/gin-gonic/gin"
	"github.com/levifleal/MyFlix/handler"
)

func initRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/serie", handler.ShowContentHandler)
		v1.POST("/serie", handler.CreateContentHandler)
		v1.DELETE("/serie", handler.DeleteContentHandler)
		v1.PUT("/serie", handler.UpdateContentHandler)
		v1.GET("/series", handler.ListContentsHandler)
	}
}
