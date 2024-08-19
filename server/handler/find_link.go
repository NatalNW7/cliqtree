package handler

import (
	"fmt"
	"net/http"

	"github.com/NatalNW7/link.in/core/schemas"
	"github.com/NatalNW7/link.in/server/helper"
	"github.com/gin-gonic/gin"
)

func FindShortLinkByLinkId(ctx *gin.Context) {
	linkID := ctx.Param("linkId")
	if linkID == "" {
		helper.SendError(ctx, http.StatusBadRequest, "linkId cannot be blank")
		return
	}

	link := schemas.Link{}
	err := db.Where("link_id = ?", linkID).First(&link).Error
	if err != nil {
		helper.SendError(ctx, http.StatusNotFound, fmt.Sprintf("Short link with linkId '%s' not found", linkID))
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, link)
}