package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateContentHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Post",
	})
}
