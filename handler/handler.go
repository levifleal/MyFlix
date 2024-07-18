package handler

import (
	"github.com/levifleal/MyFlix/handler/content"
	"github.com/levifleal/MyFlix/handler/user"
)

func InitHandler() {
	content.InitHandler()
	user.InitHandler()
}
