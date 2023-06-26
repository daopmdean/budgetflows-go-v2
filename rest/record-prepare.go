package rest

import (
	"github.com/daopmdean/budgetflows-go-v2/biz"
	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/gin-gonic/gin"
)

func PrepareIndexes(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	token := biz.GetToken(auth)
	userClaims, err := biz.ExtractToken(token)
	if err != nil {
		Response(c, UnauthorizedRes())
		return
	}

	var recordReq entity.Record
	if err := c.ShouldBindJSON(&recordReq); err != nil {
		Response(c, InvalidRes(err.Error()))
		return
	}

	Response(c, biz.PrepareIndexes(userClaims, &recordReq))
}
