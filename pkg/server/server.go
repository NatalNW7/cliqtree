package server

import "github.com/gin-gonic/gin"

func Init() error {
	server := gin.Default()
	initRoutes(server)

	err := server.Run()
	if err != nil {
		return err
	}
	
	return nil
}