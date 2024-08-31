package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/NatalNW7/cliqtree/pkg/config"
	"github.com/gin-gonic/gin"
)

func Status(ctx *gin.Context){
	env := os.Getenv("CLIQTREE_ENV")
	databaseInfo := config.DbInfo(env)
	
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"update_at": time.Now(),
		"environment": env,
		"dependencies": map[string]interface{} {
			"database": databaseInfo,
		},
	})
}