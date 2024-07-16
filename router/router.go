package router

import "github.com/gin-gonic/gin"

func Init() {
	//initialize router
	r := gin.Default()

	//initialize routes
	initRoutes(r)

	//run the server
	r.Run(":5000")
}
