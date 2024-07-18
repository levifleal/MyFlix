package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/levifleal/MyFlix/docs"
	"github.com/levifleal/MyFlix/handler"
	"github.com/levifleal/MyFlix/handler/content"
	"github.com/levifleal/MyFlix/handler/user"
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
		v1.GET("/Content", content.ShowContentHandler)
		v1.POST("/Content", content.CreateContentHandler)
		v1.DELETE("/Content", content.DeleteContentHandler)
		v1.PUT("/Content", content.UpdateContentHandler)
		v1.GET("/Contents", content.ListContentsHandler)
		v1.POST("/User", user.CreateUserHandler)
	}

	//initialize Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
