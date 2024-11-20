package rest

import (
	"github.com/daopmdean/budgetflows-go-v2/biz"
	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/daopmdean/budgetflows-go-v2/model"
	"github.com/daopmdean/summer/common"
	"github.com/gin-gonic/gin"
)

func CreateRecord(c *gin.Context) {
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

	Response(c, biz.CreateRecord(userClaims, &recordReq))
}

func UpdateRecord(c *gin.Context) {
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

	Response(c, biz.UpdateRecord(userClaims, &recordReq))
}

func GetUserRecords(c *gin.Context) {
	userClaims, err := getClaims(c)
	if err != nil {
		Response(c, common.UnauthorizedRes())
		return
	}

	var recordReq model.RecordGet
	if err := c.ShouldBindJSON(&recordReq); err != nil {
		Response(c, common.InvalidRes(err.Error()))
		return
	}

	Response(c, biz.GetUserRecords(userClaims, &recordReq))
}

func ReportUserRecords(c *gin.Context) {
	userClaims, err := getClaims(c)
	if err != nil {
		Response(c, common.UnauthorizedRes())
		return
	}

	var recordReq model.RecordReport
	if err := c.ShouldBindJSON(&recordReq); err != nil {
		Response(c, common.InvalidRes(err.Error()))
		return
	}

	Response(c, biz.ReportUserRecords(userClaims, &recordReq))
}

func DeleteUserRecord(c *gin.Context) {
	userClaims, err := getClaims(c)
	if err != nil {
		Response(c, common.UnauthorizedRes())
		return
	}

	var recordReq model.RecordDelete
	if err := c.ShouldBindJSON(&recordReq); err != nil {
		Response(c, common.InvalidRes(err.Error()))
		return
	}

	Response(c, biz.DeleteUserRecord(userClaims, &recordReq))
}
