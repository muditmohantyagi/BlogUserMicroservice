package controller

import (
	"net/http"

	"blog.com/dto"
	"blog.com/pkg/helper"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (con UserController) Register(c *gin.Context) {
	var InputDTO dto.Register

	if errDTO := c.ShouldBindJSON(&InputDTO); errDTO != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errDTO.Error())
		return
	}

	response := helper.Success(true, "ok", InputDTO)
	c.JSON(http.StatusOK, response)
	/*
		response := helper.Error("Sql Error", result.Error.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	*/
}
