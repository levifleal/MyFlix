package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteSerieHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Delete",
	})
}
