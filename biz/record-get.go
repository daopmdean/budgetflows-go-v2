package biz

import (
	"time"

	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/daopmdean/budgetflows-go-v2/model"
	"github.com/daopmdean/budgetflows-go-v2/utils"
	"github.com/daopmdean/summer/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUserRecords(userClaims *model.AppClaims, dataReq *model.RecordGet) *common.Response {
	if userClaims.UserId == 0 {
		return &common.Response{
			Status:  common.ResponseStatus.Invalid,
			Message: "UserId is required",
		}
	}

	var version string
	if dataReq.Version == "" {
		now := time.Now().In(utils.TimeZoneVN)
		version = utils.TimeToMonthlyVersion(now)
	} else {
		version = dataReq.Version
	}

	recordRes := entity.RecordDBPartition.QueryWithOpt(
		version, &entity.Record{
			UserId: userClaims.UserId,
		}, &options.FindOptions{
			Sort: &primitive.M{"record_time": -1},
		},
	)

	return recordRes
}
