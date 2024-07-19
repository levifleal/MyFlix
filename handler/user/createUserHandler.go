package user

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/levifleal/MyFlix/schemas"
	"golang.org/x/crypto/bcrypt"
)

// @BasePath /api/v1/Auth

// @Summary Create User
// @Description Create a new User
// @Tags Users
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "User data"
// @Success 200 {object} CreateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /SignUp [post]
func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	//validating user input
	if err := request.Validate(); err != nil {
		logger.Errorf("Validation Error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	//encriping password
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

	// creating user in database
	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("error creating content: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating content on database")
		return
	}

	//generating token
	token, err := newJwt(&user)
	if err != nil {
		logger.Errorf("error generating token: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error generating token")
		return
	}


	sendSuccess(ctx, "create-user", token)
}

// encript password with bcrypt
func hashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}

func newJwt(data *schemas.User) (string, error) {
	key := []byte(os.Getenv("JWT_SECRET"))

	//generating token
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  data.Name,
		"email": data.Email,
		"exp":   time.Now().Add(time.Hour * 24),
	})

	//signing token
	token, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return token, nil

}
