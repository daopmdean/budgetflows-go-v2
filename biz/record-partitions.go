package biz

import (
	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/daopmdean/summer/common"
)

func GetRecordPartitions() *common.Response {
	return &common.Response{
		Status: common.ResponseStatus.Success,
		Data: []interface{}{
			entity.RecordDBPartition.GetPartitions(),
		},
	}
}
