package biz

import (
	"time"

	"github.com/daopmdean/budgetflows-go-v2/conf"
	"github.com/daopmdean/summer/common"
)

func HealthCheck() *common.Response {
	return &common.Response{
		Status: common.ResponseStatus.Success,
		Data: []any{
			map[string]any{
				"message":   "I'm alive!",
				"startTime": conf.AppConfig.ServerStartTime,
				"timeNow":   time.Now(),
			},
		},
	}
}
