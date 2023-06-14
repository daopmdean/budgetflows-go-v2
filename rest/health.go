package rest

import (
	"net/http"

	"github.com/daopmdean/budgetflows-go-v2/biz"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, biz.HealthCheck())
}
