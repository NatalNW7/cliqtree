package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/NatalNW7/cliqtree/pkg/config"
	"github.com/NatalNW7/cliqtree/pkg/server/helper"
	"github.com/gin-gonic/gin"
)

//  @BasePath /api/v1

// @Sumary Status
// @Description Check the service status and dependencies
// @Tags Staus
// @Accept json
// @Produce json
// @Success 200 {object} helper.StatusResponse
// @Router /status [get]
func Status(ctx *gin.Context){
	env := os.Getenv("CLIQTREE_ENV")
	databaseInfo := config.DbInfo(env)
	
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, helper.StatusResponse{
		UpdatedAt: time.Now(),
		Env: env,
		Dependencies: helper.Dependencies{
			Database: databaseInfo,
		} ,
	})
}

// {
// 	"update_at": time.Now(),
// 	"environment": env,
// 	"dependencies": map[string]interface{} {
// 		"database": databaseInfo,
// 	}