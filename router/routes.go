package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/levifleal/MyFlix/docs"
	"github.com/levifleal/MyFlix/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initRoutes(r *gin.Engine) {

	//initialize handler
	handler.InitHandler()
	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	v1 := r.Group(basePath)
	{
		v1.GET("/serie", handler.ShowContentHandler)
		v1.POST("/serie", handler.CreateContentHandler)
		v1.DELETE("/serie", handler.DeleteContentHandler)
		v1.PUT("/serie", handler.UpdateContentHandler)
		v1.GET("/series", handler.ListContentsHandler)
	}

	//initialize Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
