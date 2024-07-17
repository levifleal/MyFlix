package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levifleal/MyFlix/schemas"
)

func ShowContentHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}
	content := schemas.Content{}

	if err := db.First(&content).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "content not found")
	}

	sendSuccess(ctx, "show-Content", content)
}
