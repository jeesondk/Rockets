package controllers

import (
	DTO "RocketService/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ReceiveMessage godoc
// @Summary      Receive Rocket status messages
// @Description  returns string
// @Tags         MessageService
// @Produce      plain
// @Success      200  {object}  string
// @Failure      422  {object}  string
// @Router       /messages [post]
func (c *Controller) ReceiveMessage(ctx *gin.Context) {
	var request DTO.RequestMessage
	empty := DTO.RequestMessage{}

	if err := ctx.ShouldBindJSON(&request); err != nil || request.Metadata == empty.Metadata {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid request")
	}

	err := c.MessageService.HandleMessage(request.Metadata, request.Message)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		ctx.Error(err)
	}
	ctx.JSON(http.StatusOK, "")
}
