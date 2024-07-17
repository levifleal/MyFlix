package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levifleal/MyFlix/schemas"
)

func ListContentsHandler(ctx *gin.Context) {
	contents := []schemas.Content{}

	if err := db.Find(&contents).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing contents")
	}

	sendSuccess(ctx, "list-content", contents)
}
