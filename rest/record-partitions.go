package rest

import (
	"github.com/daopmdean/budgetflows-go-v2/biz"
	"github.com/daopmdean/budgetflows-go-v2/conf"
	"github.com/daopmdean/summer/auth"
	"github.com/gin-gonic/gin"
)

func GetRecordPartitions(c *gin.Context) {
	bearer := c.Request.Header.Get("Authorization")
	token, err := auth.ExtractTokenFromHeader(bearer)
	if err != nil {
		Response(c, UnauthorizedRes())
		return
	}

	_, err = auth.ParseToken(token, conf.AppConfig.SignedKey)
	if err != nil {
		Response(c, UnauthorizedRes())
		return
	}

	Response(c, biz.GetRecordPartitions())
}
