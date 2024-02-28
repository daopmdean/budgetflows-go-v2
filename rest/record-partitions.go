package rest

import (
	"github.com/daopmdean/budgetflows-go-v2/biz"
	"github.com/gin-gonic/gin"
)

func GetRecordPartitions(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	token := biz.GetToken(auth)
	_, err := biz.ExtractToken(token)
	if err != nil {
		Response(c, UnauthorizedRes())
		return
	}

	Response(c, biz.GetRecordPartitions())
}
