package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes(server *gin.Engine){
	basePath := "/api/v1"
	routes := server.Group(basePath) 
	{
		routes.GET("/link", func(ctx *gin.Context) {
			ctx.Header("Content-type", "application/json")
			ctx.JSON(http.StatusOK, gin.H{
				"ping": "pong",
			})
		})
		routes.POST("/link", func(ctx *gin.Context) {
			ctx.Header("Content-type", "application/json")
			ctx.JSON(http.StatusCreated, gin.H{
				"ping": "pong",
			})
		})
	}
}