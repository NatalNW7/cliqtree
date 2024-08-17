package server

import (
	"net/http"

	"github.com/NatalNW7/link.in/server/handler"
	"github.com/gin-gonic/gin"
)

func initRoutes(server *gin.Engine){
	handler.Init()
	basePath := "/api/v1"
	routes := server.Group(basePath) 
	{
		routes.GET("/link", func(ctx *gin.Context) {
			ctx.Header("Content-type", "application/json")
			ctx.JSON(http.StatusOK, gin.H{
				"ping": "pong",
			})
		})
		routes.POST("/link", handler.CreateShortLink)
	}
}