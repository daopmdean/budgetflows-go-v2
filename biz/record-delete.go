package biz

import (
	"time"

	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/daopmdean/budgetflows-go-v2/model"
	"github.com/daopmdean/budgetflows-go-v2/utils"
	"github.com/daopmdean/summer/auth"
	"github.com/daopmdean/summer/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteUserRecord(userClaims *auth.SummerClaim, dataReq *model.RecordDelete) *common.Response {
	if userClaims.UserId == 0 {
		return &common.Response{
			Status:  common.ResponseStatus.Invalid,
			Message: "UserId is required",
		}
	}

	if dataReq.Id == "" {
		return &common.Response{
			Status:  common.ResponseStatus.Invalid,
			Message: "Record Id is required",
		}
	}

	var version string
	if dataReq.Version == "" {
		now := time.Now().In(utils.TimeZoneVN)
		version = utils.TimeToMonthlyVersion(now)
	} else {
		version = dataReq.Version
	}

	objId, err := primitive.ObjectIDFromHex(dataReq.Id)
	if err != nil {
		return &common.Response{
			Status:  common.ResponseStatus.Invalid,
			Message: "Invalid Record Id",
		}
	}

	deleteRes := entity.RecordDBPartition.Delete(version, &entity.Record{
		ID: &objId,
	})

	return deleteRes
}
