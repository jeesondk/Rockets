package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) Messages(ctx *gin.Context) {
	ctx.JSON(http.StatusBadGateway, "")
}
