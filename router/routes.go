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
	
	//defining BasePath for V1
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := r.Group(basePath)

	//defining Auth routes --Unprotected
	auth := v1.Group("/Auth")
	{
		auth.POST("/SignUp", user.CreateUserHandler)
		auth.POST("/SignIn", user.LoginUserHandler)

	}

	//defininig routes --Protected
	p := v1.Group("/").Use(AuthMiddleware())
	{
		p.GET("/Content", content.ShowContentHandler)
		p.POST("/Content", content.CreateContentHandler)
		p.DELETE("/Content", content.DeleteContentHandler)
		p.PUT("/Content", content.UpdateContentHandler)
		p.GET("/Contents", content.ListContentsHandler)
	}

	//initialize Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
