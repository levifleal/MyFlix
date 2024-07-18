package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levifleal/MyFlix/schemas"
	"golang.org/x/crypto/bcrypt"
)

// @BasePath /api/v1

// @Summary Create User
// @Description Create a new User
// @Tags Users
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "User data"
// @Success 200 {object} CreateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /User [post]
func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation Error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	hashPass, err := hashPassword(request.Password)
	if err != nil {
		logger.Error("error crypting password")
		sendError(ctx, http.StatusInternalServerError, "error crypting password")
		return
	}

	user := schemas.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashPass,
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("error creating content: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating content on database")
		return
	}

	sendSuccess(ctx, "Create-Content", user)

}

func hashPassword(pass string) (hashPass string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}
