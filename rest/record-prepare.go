package rest

import (
	"github.com/daopmdean/budgetflows-go-v2/biz"
	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/daopmdean/summer/common"
	"github.com/gin-gonic/gin"
)

func PrepareIndexes(c *gin.Context) {
	userClaims, err := getClaims(c)
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
