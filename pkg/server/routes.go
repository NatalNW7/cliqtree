package server

import (
	"github.com/NatalNW7/cliqtree/docs"
	"github.com/NatalNW7/cliqtree/pkg/server/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initRoutes(server *gin.Engine){
	handler.Init()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	routes := server.Group(basePath) 
	{
		routes.GET("/link/:linkId", handler.FindShortLinkByLinkId)
		routes.POST("/link", handler.CreateShortLink)
		routes.GET("/redirect/:linkId", handler.RedirectUrl)
		routes.GET("/status", handler.Status)
		routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

}