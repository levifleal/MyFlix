package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSerieHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Post",
	})
}
func ShowSerieHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Get",
	})
}
func DeleteSerieHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Delete",
	})
}
func UpdateSerieHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Put",
	})
}
func ListSeriesHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Get All",
	})
}
