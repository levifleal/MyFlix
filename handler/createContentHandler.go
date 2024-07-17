package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levifleal/MyFlix/schemas"
)

// @BasePath /api/v1

// @Summary Create Content
// @Description Create a new Media Content
// @Tags Media Contents
// @Accept json
// @Produce json
// @Param request body CreateContentRequest true "Content data"
// @Success 200 {object} CreateContentResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /Content [post]
func CreateContentHandler(ctx *gin.Context) {
	request := CreateContentRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation Error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	content := schemas.Content{
		Title:       request.Title,
		Desc:        request.Desc,
		ContentType: request.ContentType,
		Genre:       request.Genre,
		ReleaseDate: request.ReleaseDate,
	}

	if err := db.Create(&content).Error; err != nil {
		logger.Errorf("error creating content: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating content on database")
		return
	}

	sendSuccess(ctx, "Create-Content", content)

}
