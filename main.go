package main

import (
	"github.com/levifleal/MyFlix/config"
	"github.com/levifleal/MyFlix/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	//initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//initializer router
	router.Init()
}
