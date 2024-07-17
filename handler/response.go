package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levifleal/MyFlix/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("content-type", "application/json")
	ctx.JSON(code, gin.H{
		"msg": msg,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  fmt.Sprintf("success operation from handler: %s", op),
		"data": data,
	})
}

type ErrorResponse struct {
	Message string `json:"msg"`
}
type CreateContentResponse struct {
	Message string                  `json:"msg"`
	Data    schemas.ContentResponse `json:"data"`
}
type DeleteContentResponse struct {
	Message string                  `json:"msg"`
	Data    schemas.ContentResponse `json:"data"`
}
type ShowContentResponse struct {
	Message string                  `json:"msg"`
	Data    schemas.ContentResponse `json:"data"`
}
type ListContentsResponse struct {
	Message string                    `json:"msg"`
	Data    []schemas.ContentResponse `json:"data"`
}
type UpdateContentsResponse struct {
	Message string                    `json:"msg"`
	Data    []schemas.ContentResponse `json:"data"`
}
