package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllRockets godoc
// @Summary      Returns all rockets
// @Description  returns array of rockets
// @Tags         RocketService
// @Produce      json
// @Success      200  {object}  []entities.Rocket
// @Failure      500  {object}  string
// @Router       /rockets [get]
func (c *Controller) GetAllRockets(ctx *gin.Context) {
	res, err := c.RocketService.GetAllRockets()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(200, res)
}

// GetRocket godoc
// @Summary      Return a rocket by id
// @Description  returns instance of rocket
// @Tags         RocketService
// @Param   	 id     query    string     true        "id"
// @Produce      json
// @Success      200  {object}  entities.Rocket
// @Failure      500  {object}  string
// @Router       /rocket [get]
func (c *Controller) GetRocket(ctx *gin.Context) {
	id := ctx.Query("id")
	res, err := c.RocketService.GetRocket(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}
	ctx.JSON(200, res)
}
