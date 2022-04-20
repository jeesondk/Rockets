package controllers

import (
	DTO "RocketService/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) RecieveMessage(ctx *gin.Context) {
	var request DTO.RequestMessage
	empty := DTO.RequestMessage{}

	if err := ctx.ShouldBindJSON(&request); err != nil || request.Metadata == empty.Metadata {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid request")
		return
	}

	ctx.JSON(http.StatusOK, "")
}
