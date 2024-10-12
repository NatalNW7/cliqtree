package handler

import (
	"fmt"
	"net/http"

	"github.com/NatalNW7/cliqtree/internal/cliqtree/usecases"
	"github.com/NatalNW7/cliqtree/pkg/server/helper"
	"github.com/gin-gonic/gin"
)

//  @BasePath /api/v1

// @Sumary Redirect Url
// @Description Find short link by link id and redirect user to redirectUrl
// @Tags Links
// @Accept json
// @Produce json
// @Param linkId path int true "linkID"
// @Success 301
// @Failure 400 {object} helper.ErrorResponse
// @Failure 404 {object} helper.ErrorResponse
// @Failure 405 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /redirect/{linkId} [get]
func RedirectUrl(ctx *gin.Context) {
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

	ctx.Header("Location", link.RedirectUrl)
	ctx.Redirect(http.StatusMovedPermanently, link.RedirectUrl)
}