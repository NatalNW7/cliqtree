package handler

import (
	"net/http"

	"github.com/NatalNW7/link.in/internal/link.in/usecases"
	"github.com/NatalNW7/link.in/pkg/server/helper"
	"github.com/gin-gonic/gin"
)


func CreateShortLink(ctx *gin.Context) {
	body := helper.LinkRequest{}

	ctx.BindJSON(&body)
	err := body.Validade()
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
