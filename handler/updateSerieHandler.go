package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateSerieHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Put",
	})
}
