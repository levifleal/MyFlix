package content

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levifleal/MyFlix/schemas"
)

// @BasePath /api/v1

// @Summary Show Content
// @Description Show a Media Content
// @Tags Media Contents
// @Accept json
// @Produce json
// @Param id query string true "Content Identification"
// @Success 200 {object} ShowContentResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /Content [get]
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
