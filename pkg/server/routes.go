package server

import (
	"github.com/NatalNW7/link.in/pkg/server/handler"
	"github.com/gin-gonic/gin"
)

func initRoutes(server *gin.Engine){
	handler.Init()
	basePath := "/api/v1"
	routes := server.Group(basePath) 
	{
		routes.GET("/link/:linkId", handler.FindShortLinkByLinkId)
		routes.POST("/link", handler.CreateShortLink)
		routes.GET("/redirect/:linkId", handler.RedirectUrl)
		routes.GET("/status", handler.Status)
	}
}