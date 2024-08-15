package handler

import (
	"net/http"

	"github.com/NatalNW7/link.in/core/schemas"
	"github.com/NatalNW7/link.in/core/util"
	"github.com/NatalNW7/link.in/server/helper"
	"github.com/gin-gonic/gin"
)


func CreateShortLink(ctx *gin.Context) {
	body := helper.LinkRequest{}

	ctx.BindJSON(&body)
	err := body.Validade()
	if err != nil {
		logger.Errorf("Error to validate body request: %v", err.Error())
		helper.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	link := schemas.Link{
		LinkId: util.GenerateLinkId(),
		RedirectUrl: body.Url,
	}

	err = db.Create(&link).Error
	if err != nil {
		logger.Errorf("Error in saving shortened link on the database: %v", err.Error())
		helper.SendError(ctx, http.StatusInternalServerError, "Error in saving shortened link on the database")
		return
	}

	helper.SendSuccess(ctx, http.StatusCreated, link)
}
