package rest

import (
	"github.com/daopmdean/budgetflows-go-v2/biz"
	"github.com/daopmdean/budgetflows-go-v2/conf"
	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/daopmdean/summer/auth"
	"github.com/daopmdean/summer/common"
	"github.com/gin-gonic/gin"
)

func PrepareIndexes(c *gin.Context) {
	bearer := c.Request.Header.Get("Authorization")
	token, err := auth.ExtractTokenFromHeader(bearer)
	if err != nil {
		Response(c, common.UnauthorizedRes())
		return
	}

	userClaims, err := auth.ParseToken(token, conf.AppConfig.SignedKey)
	if err != nil {
		Response(c, common.UnauthorizedRes())
		return
	}

	var recordReq entity.Record
	if err := c.ShouldBindJSON(&recordReq); err != nil {
		Response(c, common.InvalidRes(err.Error()))
		return
	}

	Response(c, biz.PrepareIndexes(userClaims, &recordReq))
}
