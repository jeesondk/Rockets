package controllers

import (
	DTO "RocketService/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ReceiveMessage godoc
// @Summary      Receive Rocket status messages
// @Description  returns string
// @Tags         Message Service
// @Produce      plain
// @Success      200  {object}  string
// @Router       /messages [post]
func (c *Controller) ReceiveMessage(ctx *gin.Context) {
	var request DTO.RequestMessage
	empty := DTO.RequestMessage{}

	if err := ctx.ShouldBindJSON(&request); err != nil || request.Metadata == empty.Metadata {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid request")
		return
	}

	ctx.JSON(http.StatusOK, "")
}
