package server

import "github.com/gin-gonic/gin"

func Init() {
	server := gin.Default()
	initRoutes(server)
	server.Run()
}