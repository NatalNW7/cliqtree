package handler

import (
	"fmt"
	"net/http"

	"github.com/NatalNW7/link.in/internal/link.in/usecases"
	"github.com/NatalNW7/link.in/pkg/server/helper"
	"github.com/gin-gonic/gin"
)

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