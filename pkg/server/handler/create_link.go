package handler

import (
	"net/http"

	"github.com/NatalNW7/cliqtree/internal/cliqtree/usecases"
	"github.com/NatalNW7/cliqtree/pkg/server/helper"
	"github.com/gin-gonic/gin"
)

//  @BasePath /api/v1

// @Sumary Create Short Link
// @Description Create a new short link
// @Tags Links
// @Accept json
// @Produce json
// @Param request body helper.LinkRequest true "Request body"
// @Success 201 {object} helper.SuccessResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 405 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /link [post]
func CreateShortLink(ctx *gin.Context) {
	err := helper.MethodAllowed(ctx.Request.Method)
	if err != nil {
		helper.SendError(ctx, http.StatusMethodNotAllowed, err.Error(), err.Error())
		return		
	}

	body := helper.LinkRequest{}

	err = ctx.BindJSON(&body)
	if err != nil {
		logger.Errorf("Error to bind body request: %v", err.Error())
		helper.SendError(ctx, http.StatusInternalServerError, "Error to body body request", err.Error())
		return
	}
	
	err = body.Validade()
	if err != nil {
		logger.Errorf("Error to validate body request: %v", err.Error())
		helper.SendError(ctx, http.StatusBadRequest, "Error to validate body request", err.Error())
		return
	}

	link, err := usecases.CreateShortLink(body.Url, db, logger)
	if err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, "Error to save shortened link on database", err.Error())
		return
	}



	helper.SendSuccess(ctx, http.StatusCreated, link)
}
