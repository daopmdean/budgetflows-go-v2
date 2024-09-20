package rest

import (
	"github.com/daopmdean/budgetflows-go-v2/biz"
	"github.com/daopmdean/summer/common"
	"github.com/gin-gonic/gin"
)

func GetRecordPartitions(c *gin.Context) {
	_, err := getClaims(c)
	if err != nil {
		Response(c, common.UnauthorizedRes())
		return
	}

	Response(c, biz.GetRecordPartitions())
}
