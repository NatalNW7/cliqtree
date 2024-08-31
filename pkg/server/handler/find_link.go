package handler

import (
	"fmt"
	"net/http"

	"github.com/NatalNW7/cliqtree/internal/cliqtree/usecases"
	"github.com/NatalNW7/cliqtree/pkg/server/helper"
	"github.com/gin-gonic/gin"
)

//  @BasePath /api/v1

// @Sumary Find Link
// @Description Find short link by link id
// @Tags Links
// @Accept json
// @Produce json
// @Param linkId path int true "linkID"
// @Success 200 {object} helper.SuccessResponse
// @Failure 400 {object} helper.ErrorResponse
// @Failure 405 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /link/{linkId} [get]
func FindShortLinkByLinkId(ctx *gin.Context) {
	err := helper.MethodAllowed(ctx.Request.Method)
	if err != nil {
		helper.SendError(ctx, http.StatusMethodNotAllowed, err.Error(), err.Error())
		return		
	}

	linkID, err := helper.ValidateLinkId(ctx.Param("linkId"))
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "linkId cannot be blank", err.Error())
		return
	}

	link, err := usecases.FindShortLinkByLinkId(linkID, db, logger)
	if err != nil {
		helper.SendError(ctx, http.StatusNotFound, fmt.Sprintf("Short link with linkId '%s' not found", linkID), err.Error())
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, link)
}