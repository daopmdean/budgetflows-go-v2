package biz

import (
	"github.com/daopmdean/budgetflows-go-v2/model"
	"github.com/daopmdean/summer/common"
	"github.com/gin-gonic/gin"
)

func Register(data *model.RegisterRequest) *common.Response {
	return &common.Response{
		Status: common.ResponseStatus.Success,
		Data: []any{
			gin.H{
				"message": "Register success!",
			},
		},
	}
}
