package main

import (
	"github.com/NatalNW7/link.in/pkg/config"
	"github.com/NatalNW7/link.in/pkg/server"
)

func main(){
	logger := config.GetLogger("root")
	err := config.Init()
	if err != nil {
		logger.Errorf("Error to initialize project configuration: %v", err)
		return
	}
	
	err = server.Init()
	if err != nil {
		logger.Errorf("Error to initialize server: %v", err)
		return
	}
}