package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteContentHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Delete",
	})
}
