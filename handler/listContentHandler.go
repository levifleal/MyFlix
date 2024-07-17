package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levifleal/MyFlix/schemas"
)

// @BasePath /api/v1

// @Summary List Contents
// @Description list all Media Contents
// @Tags Media Contents
// @Accept json
// @Produce json
// @Success 200 {object} ListContentsResponse
// @Failure 500 {object} ErrorResponse
// @Router /Contents [get]
func ListContentsHandler(ctx *gin.Context) {
	contents := []schemas.Content{}

	if err := db.Find(&contents).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing contents")
	}

	sendSuccess(ctx, "list-content", contents)
}
