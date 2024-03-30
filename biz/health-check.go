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
				"message":     "I'm alive!",
				"startedTime": conf.AppConfig.ServerStartTime,
				"appName":     conf.AppConfig.AppName,
				"env":         conf.AppConfig.Env,
				"timeNow":     time.Now(),
				"flag":        "version-1.0.0",
			},
		},
	}
}
