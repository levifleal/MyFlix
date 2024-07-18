package content

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levifleal/MyFlix/schemas"
)

// @BasePath /api/v1

// @Summary Delete Content
// @Description Delete a new Media Content
// @Tags Media Contents
// @Accept json
// @Produce json
// @Param id query string true "Content Identification"
// @Success 200 {object} DeleteContentResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /Content [delete]
func DeleteContentHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}
	content := schemas.Content{}
	//finding content
	if err := db.First(&content, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("content with id: %s not found", id))
		return
	}
	//deleting contentent
	if err := db.Delete(&content).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting content with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-opening", content)
}
