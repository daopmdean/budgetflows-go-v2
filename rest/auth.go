package rest

import (
	"net/http"

	"github.com/daopmdean/budgetflows-go-v2/biz"
	"github.com/daopmdean/budgetflows-go-v2/model"
	"github.com/daopmdean/budgetflows-go-v2/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginReq model.LoginRequest
	if err := c.BindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.Invalid(err.Error()))
		return
	}

	c.JSON(http.StatusOK, biz.Login(&loginReq))
}

func Register(c *gin.Context) {
	var registerReq model.RegisterRequest
	if err := c.BindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.Invalid(err.Error()))
		return
	}

	c.JSON(http.StatusOK, biz.Register(&registerReq))
}
