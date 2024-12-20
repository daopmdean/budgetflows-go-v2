package biz

import (
	"time"

	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/daopmdean/budgetflows-go-v2/model"
	"github.com/daopmdean/budgetflows-go-v2/utils"
	"github.com/daopmdean/summer/auth"
	"github.com/daopmdean/summer/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReportUserRecords(
	userClaims *auth.SummerClaim,
	dataReq *model.RecordReport,
) *common.Response {
	if userClaims.UserId == 0 {
		return common.BuildInvalidRes("USER_ID_REQUIRED", "UserId is required")
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
	if recordRes.Status != common.ResponseStatus.Success {
		return recordRes
	}

	totalAmount := 0.0

	tagsMap := map[string]float64{}
	records := recordRes.Data.([]*entity.Record)
	for _, record := range records {
		totalAmount += record.Amount
		for _, tag := range record.Tags {
			tagsMap[tag] += record.Amount
		}
	}

	return &common.Response{
		Status: common.ResponseStatus.Success,
		Data: []any{
			map[string]any{
				"timeKey":     version,
				"totalAmount": totalAmount,
				"tags":        tagsMap,
			},
		},
	}
}
