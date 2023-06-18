package rest

import (
	"github.com/daopmdean/budgetflows-go-v2/biz"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	Response(c, biz.HealthCheck())
}
