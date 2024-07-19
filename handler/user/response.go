package user

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
type CreateUserResponse struct {
	Message string               `json:"msg"`
	Data    schemas.UserResponse `json:"data"`
}
type LoginUserResponse struct {
	Message string               `json:"msg"`
	Data    schemas.UserResponse `json:"data"`
}
