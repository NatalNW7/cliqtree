package main

import (
	"github.com/NatalNW7/link.in/config"
	"github.com/NatalNW7/link.in/server"
)

func main(){
	logger := config.GetLogger("root")
	err := config.Init()
	if err != nil {
		logger.Errorf("Error to initialize project configuration: %v", err)
		return
	}

	server.Init()
}