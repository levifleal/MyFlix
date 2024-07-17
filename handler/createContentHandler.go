package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levifleal/MyFlix/schemas"
)

func CreateContentHandler(ctx *gin.Context) {
	request := CreateContentRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation Error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	content := schemas.Content{
		Title:        request.Title,
		Desc:         request.Desc,
		ContentType:  request.ContentType,
		Genre:        request.Genre,
		RealeaseDate: request.ReleaseDate,
	}

	if err := db.Create(&content).Error; err != nil {
		logger.Errorf("error creating content: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating content on database")
		return
	} 

	sendSuccess(ctx,"Create-Content",content)

}
