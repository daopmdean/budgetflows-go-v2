package rest

import (
	"github.com/daopmdean/budgetflows-go-v2/biz"
	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/daopmdean/budgetflows-go-v2/model"
	"github.com/gin-gonic/gin"
)

func CreateRecord(c *gin.Context) {
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

	Response(c, biz.CreateRecord(userClaims, &recordReq))
}

func GetUserRecords(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	token := biz.GetToken(auth)
	userClaims, err := biz.ExtractToken(token)
	if err != nil {
		Response(c, UnauthorizedRes())
		return
	}

	var recordReq model.RecordGet
	if err := c.ShouldBindJSON(&recordReq); err != nil {
		Response(c, InvalidRes(err.Error()))
		return
	}

	Response(c, biz.GetUserRecords(userClaims, &recordReq))
}

func DeleteUserRecord(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	token := biz.GetToken(auth)
	userClaims, err := biz.ExtractToken(token)
	if err != nil {
		Response(c, UnauthorizedRes())
		return
	}

	var recordReq model.RecordDelete
	if err := c.ShouldBindJSON(&recordReq); err != nil {
		Response(c, InvalidRes(err.Error()))
		return
	}

	Response(c, biz.DeleteUserRecord(userClaims, &recordReq))
}
