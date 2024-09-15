package rest

import (
	"github.com/daopmdean/budgetflows-go-v2/biz"
	"github.com/daopmdean/budgetflows-go-v2/model"
	"github.com/daopmdean/summer/common"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginReq model.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		Response(c, common.InvalidRes(err.Error()))
		return
	}

	Response(c, biz.Login(&loginReq))
}

func Register(c *gin.Context) {
	var registerReq model.RegisterRequest
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		Response(c, common.InvalidRes(err.Error()))
		return
	}

	Response(c, biz.Register(&registerReq))
}
