package rest

import (
	"net/http"
	"time"

	"github.com/daopmdean/budgetflows-go-v2/conf"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":   "I'm alive!",
		"startTime": conf.AppConfig.ServerStartTime,
		"timeNow":   time.Now(),
	})
}
