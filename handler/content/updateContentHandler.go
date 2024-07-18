package content

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levifleal/MyFlix/schemas"
)

// @BasePath /api/v1

// @Summary Update Content
// @Description Update a Media Content
// @Tags Media Contents
// @Accept json
// @Produce json
// @Param id query string true "Content Identification"
// @Param request body updateContentRequest true "Content data"
// @Success 200 {object} UpdateContentsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /Content [put]
func UpdateContentHandler(ctx *gin.Context) {
	request := updateContentRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}

	content := schemas.Content{}

	if err := db.First(&content, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "content not found")
		return
	}

	if request.Title != "" {
		content.Title = request.Title
	}
	if request.Desc != "" {
		content.Desc = request.Desc
	}
	if request.Genre != "" {
		content.Genre = request.Genre
	}
	if request.ContentType != "" {
		content.ContentType = request.ContentType
	}
	if !request.ReleaseDate.IsZero() {
		content.ReleaseDate = request.ReleaseDate
	}

	if err := db.Save(&content).Error; err != nil {
		logger.Errorf("error updating content: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating content")
		return
	}

	sendSuccess(ctx, "update-content", content)

}
