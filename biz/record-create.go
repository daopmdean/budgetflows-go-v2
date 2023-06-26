package biz

import (
	"time"

	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/daopmdean/budgetflows-go-v2/model"
	"github.com/daopmdean/budgetflows-go-v2/utils"
	"github.com/daopmdean/summer/common"
)

func CreateRecord(userClaims *model.AppClaims, data *entity.Record) *common.Response {
	if userClaims.UserId == 0 {
		return &common.Response{
			Status:  common.ResponseStatus.Invalid,
			Message: "UserId is required",
		}
	}

	data.UserId = userClaims.UserId

	if data.RecordTime == nil {
		now := time.Now().In(utils.TimeZoneVN)
		data.RecordTime = &now
	}

	version := utils.TimeToMonthlyVersion(*data.RecordTime)
	recordRes := entity.RecordDBPartition.Create(version, data)
	if recordRes.Status != common.ResponseStatus.Success {
		return recordRes
	}

	return recordRes
}
