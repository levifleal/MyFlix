package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/levifleal/MyFlix/config"
)

var (
	logger *config.Logger
)

func Init() {

	//initialize logger
	logger = config.GetLogger("[ Router ]")

	//initialize router
	r := gin.Default()

	//initialize routes
	initRoutes(r)

	//initializing middlewares
	r.Use(cors.Default())

	//run the server
	r.Run(":5000")
}
