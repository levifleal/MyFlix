package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/serie", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "Get",
			})
		})
		v1.POST("/serie", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "Post",
			})
		})
		v1.DELETE("/serie", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "Delete",
			})
		})
		v1.PUT("/serie", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "Put",
			})
		})
		v1.GET("/series", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "Gets",
			})
		})
	}
}
