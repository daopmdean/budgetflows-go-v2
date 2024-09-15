package rest

import (
	"github.com/daopmdean/budgetflows-go-v2/biz"
	"github.com/daopmdean/budgetflows-go-v2/conf"
	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/daopmdean/budgetflows-go-v2/model"
	"github.com/daopmdean/summer/auth"
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

func getClaims(c *gin.Context) (*auth.SummerClaim, error) {
	bearer := c.Request.Header.Get("Authorization")
	token, err := auth.ExtractTokenFromHeader(bearer)
	if err != nil {
		return nil, err
	}

	return auth.ParseToken(token, conf.AppConfig.SignedKey)
}

func UpdateRecord(c *gin.Context) {
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

	Response(c, biz.UpdateRecord(userClaims, &recordReq))
}

func GetUserRecords(c *gin.Context) {
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

	var recordReq model.RecordGet
	if err := c.ShouldBindJSON(&recordReq); err != nil {
		Response(c, common.InvalidRes(err.Error()))
		return
	}

	Response(c, biz.GetUserRecords(userClaims, &recordReq))
}

func ReportUserRecords(c *gin.Context) {
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

	var recordReq model.RecordGet
	if err := c.ShouldBindJSON(&recordReq); err != nil {
		Response(c, common.InvalidRes(err.Error()))
		return
	}

	Response(c, biz.ReportUserRecords(userClaims, &recordReq))
}

func DeleteUserRecord(c *gin.Context) {
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

	var recordReq model.RecordDelete
	if err := c.ShouldBindJSON(&recordReq); err != nil {
		Response(c, common.InvalidRes(err.Error()))
		return
	}

	Response(c, biz.DeleteUserRecord(userClaims, &recordReq))
}
