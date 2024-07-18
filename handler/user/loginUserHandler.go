package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levifleal/MyFlix/schemas"
	"golang.org/x/crypto/bcrypt"
)

func LoginUserHandler(ctx *gin.Context) {
	request := LoginUserRequest{}

	ctx.BindJSON(&request)

	//validating user input
	if err := request.Validate(); err != nil {
		logger.Errorf("Validation Error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	//finding user in database
	user := schemas.User{}
	if err := db.First(&user, schemas.User{Email: request.Email}).Error; err != nil {
		sendError(ctx, http.StatusUnauthorized, "invalid credentials")
		return
	}

	//checking if passwords match
	passMatch := CheckPasswordHash(request.Password, user.Password)
	if !passMatch {
		sendError(ctx, http.StatusUnauthorized, "invalid credentials")
		return
	}

	//generating token
	token, err := newJwt(&user)
	if err != nil {
		logger.Errorf("error generating token: %v", err.Error())
		return
	}

	sendSuccess(ctx, "login-user", token)
}

// checking if passwords match thru bcrypt
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
