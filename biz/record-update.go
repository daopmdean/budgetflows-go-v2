package biz

import (
	"time"

	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/daopmdean/budgetflows-go-v2/utils"
	"github.com/daopmdean/summer/auth"
	"github.com/daopmdean/summer/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateRecord(userClaims *auth.SummerClaim, data *entity.Record) *common.Response {
	if userClaims.UserId == 0 {
		return &common.Response{
			Status:  common.ResponseStatus.Invalid,
			Message: "UserId is required",
		}
	}

	if data.ID == nil || *data.ID == primitive.NilObjectID {
		return &common.Response{
			Status:  common.ResponseStatus.Invalid,
			Message: "Record Id is required",
		}
	}

	var version string
	if data.Version == "" {
		now := time.Now().In(utils.TimeZoneVN)
		version = utils.TimeToMonthlyVersion(now)
	} else {
		version = data.Version
	}

	existRecordRes := entity.RecordDBPartition.QueryOne(version, &entity.Record{
		ID: data.ID,
	})
	if existRecordRes.Status != common.ResponseStatus.Success {
		return existRecordRes
	}

	existRecord := existRecordRes.Data.([]*entity.Record)[0]
	if existRecord.UserId != userClaims.UserId {
		return &common.Response{
			Status:  common.ResponseStatus.Unauthorized,
			Message: "Not authorized to update record",
		}
	}

	existedVersion := utils.TimeToMonthlyVersion(*existRecord.RecordTime)

	if data.RecordTime != nil {
		version := utils.TimeToMonthlyVersion(*data.RecordTime)
		if existedVersion != version {
			return &common.Response{
				Status:  common.ResponseStatus.Invalid,
				Message: "Cannot update record with different version",
			}
		}
	}

	recordRes := entity.RecordDBPartition.Upsert(existedVersion, &entity.Record{
		ID: data.ID,
	}, data)
	if recordRes.Status != common.ResponseStatus.Success {
		return recordRes
	}

	return recordRes
}
