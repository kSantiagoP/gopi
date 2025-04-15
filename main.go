package main

import (
	"fmt"

	"github.com/kSantiagoP/gopi/config"
	"github.com/kSantiagoP/gopi/router"
)

func main() {
	//Initialize configs
	err := config.Init()

	if err != nil {
		fmt.Println("Error initializing configs")
		return
	}

	//Initialize router
	router.Initialize()
}
