package biz

import (
	"fmt"
	"time"

	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/daopmdean/budgetflows-go-v2/utils"
	"github.com/daopmdean/summer/auth"
	"github.com/daopmdean/summer/common"
)

func PrepareIndexes(userClaims *auth.SummerClaim, data *entity.Record) *common.Response {
	if userClaims.UserId == 0 {
		return &common.Response{
			Status:  common.ResponseStatus.Invalid,
			Message: "UserId is required",
		}
	}

	var version string
	if data.Version == "" {
		now := time.Now().In(utils.TimeZoneVN)
		version = utils.TimeToMonthlyVersion(now)
	} else {
		version = data.Version
	}

	err := entity.RecordDBPartition.PrepareCol(version)
	if err != nil {
		return common.BuildErrorRes("PREPARE_ERROR", err.Error())
	}

	return &common.Response{
		Status:  common.ResponseStatus.Success,
		Message: fmt.Sprintf("create index with version %s success", version),
	}
}
