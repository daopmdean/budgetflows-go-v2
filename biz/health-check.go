package biz

import (
	"time"

	"github.com/daopmdean/budgetflows-go-v2/conf"
	"github.com/daopmdean/summer/common"
	"github.com/gin-gonic/gin"
)

func HealthCheck() *common.Response {
	return &common.Response{
		Status: common.ResponseStatus.Success,
		Data: []any{
			gin.H{
				"message":   "I'm alive!",
				"startTime": conf.AppConfig.ServerStartTime,
				"timeNow":   time.Now(),
			},
		},
	}
}
