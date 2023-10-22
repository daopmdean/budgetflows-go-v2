package biz

import (
	"github.com/daopmdean/budgetflows-go-v2/model"
	"github.com/daopmdean/summer/common"
)

func GetRecentUserRecords(userClaims *model.AppClaims) *common.Response {
	return GetUserRecords(userClaims, &model.RecordGet{})
}
