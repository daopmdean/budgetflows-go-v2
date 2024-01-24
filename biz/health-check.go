package biz

import (
	"time"

	"github.com/daopmdean/budgetflows-go-v2/conf"
	"github.com/daopmdean/budgetflows-go-v2/model"
	"github.com/daopmdean/summer/common"
)

func HealthCheck() *common.Response {
	return &common.Response{
		Status: common.ResponseStatus.Success,
		Data: []any{
			model.HealthStatus{
				Message:     "I'm alive!",
				StartedTime: conf.AppConfig.ServerStartTime,
				AppName:     "budgetflows-go-v2",
				TimeNow:     time.Now(),
			},
		},
	}
}
