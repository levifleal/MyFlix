package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListSeriesHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Get All",
	})
}
