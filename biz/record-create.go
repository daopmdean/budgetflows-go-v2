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

	now := time.Now().In(utils.TimeZoneVN)
	version := utils.TimeToKey(now)

	data.UserId = userClaims.UserId

	if data.RecordTime == nil {
		data.RecordTime = &now
	}

	recordRes := entity.RecordDBPartition.Create(version, data)
	if recordRes.Status != common.ResponseStatus.Success {
		return recordRes
	}

	return recordRes
}
