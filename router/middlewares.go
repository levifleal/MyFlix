package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//verifing if has Autorization key
		bearerToken := ctx.Request.Header.Get("Authorization")
		token, _ := strings.CutPrefix(bearerToken, "Bearer ")
		if token == "" {
			ctx.Header("content-type", "application/json")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "token is missing",
			})
			return
		}
		validator := jwt.NewParser(jwt.WithExpirationRequired())
		_, _, err := validator.ParseUnverified(token, jwt.MapClaims{})
		if err != nil {
			ctx.Header("content-type", "application/json")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": fmt.Sprint(err),
			})
			return
		}
	}
}
