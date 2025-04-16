package main

import (
	"github.com/kSantiagoP/gopi/config"
	"github.com/kSantiagoP/gopi/router"
)

var (
	logger *config.Logger
)

func main() {

	//Initialize configs
	logger = config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//Initialize router
	router.Initialize()
}
