package handler

import (
	"fmt"
	"net/http"

	"github.com/NatalNW7/link.in/internal/link.in/usecases"
	"github.com/NatalNW7/link.in/pkg/server/helper"
	"github.com/gin-gonic/gin"
)

func FindShortLinkByLinkId(ctx *gin.Context) {
	linkID := ctx.Param("linkId")
	if linkID == "" {
		helper.SendError(ctx, http.StatusBadRequest, "linkId cannot be blank", "")
		return
	}

	link, err := usecases.FindShortLinkByLinkId(linkID, db, logger)
	if err != nil {
		helper.SendError(ctx, http.StatusNotFound, fmt.Sprintf("Short link with linkId '%s' not found", linkID), err.Error())
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, link)
}