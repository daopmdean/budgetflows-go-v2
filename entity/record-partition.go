package entity

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/daopmdean/summer/common"
	"github.com/daopmdean/summer/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MonthlyPartitionInstance struct {
	mongodb.Instance
	database *mongo.Database
}

var RecordDBPartition = &MonthlyPartitionInstance{
	Instance: mongodb.Instance{
		ColName:     "record",
		TemplateObj: new(Record),
	},
}

func (pi *MonthlyPartitionInstance) SetDBPartition(
	database *mongo.Database,
) *MonthlyPartitionInstance {
	pi.Instance.SetDB(database)
	pi.database = database
	return pi
}

func (pi *MonthlyPartitionInstance) Create(
	version string,
	ele interface{},
) *common.Response {
	if !validPartitionVersion(version) {
		return &common.Response{
			Status: common.ResponseStatus.Error,
			Error: &common.ErrorResponse{
				ErrorCode:    "INVALID_PARTITION_VERSION",
				ErrorMessage: "Invalid partition version: " + version,
			},
		}
	}

	pi.ColName = getCollectionName(version)
	pi.Instance.SetDB(pi.database)
	return pi.Instance.Create(context.TODO(), ele)
}

func (pi *MonthlyPartitionInstance) Query(
	version string,
	filter interface{},
	skip, limit int64,
	sortFields *primitive.M,
) *common.Response {
	if !validPartitionVersion(version) {
		return &common.Response{
			Status: common.ResponseStatus.Error,
			Error: &common.ErrorResponse{
				ErrorCode:    "INVALID_PARTITION_VERSION",
				ErrorMessage: "Invalid partition version: " + version,
			},
		}
	}

	pi.ColName = getCollectionName(version)
	pi.Instance.SetDB(pi.database)
	return pi.Instance.Query(context.TODO(), filter, skip, limit, sortFields)
}

func (pi *MonthlyPartitionInstance) QueryOne(
	version string,
	filter interface{},
) *common.Response {
	if !validPartitionVersion(version) {
		return &common.Response{
			Status: common.ResponseStatus.Error,
			Error: &common.ErrorResponse{
				ErrorCode:    "INVALID_PARTITION_VERSION",
				ErrorMessage: "Invalid partition version: " + version,
			},
		}
	}

	pi.ColName = getCollectionName(version)
	pi.Instance.SetDB(pi.database)
	return pi.Instance.QueryOne(context.TODO(), filter)
}

func (pi *MonthlyPartitionInstance) QueryWithOpt(
	version string,
	filter interface{},
	opt *options.FindOptions,
) *common.Response {
	if !validPartitionVersion(version) {
		return &common.Response{
			Status: common.ResponseStatus.Error,
			Error: &common.ErrorResponse{
				ErrorCode:    "INVALID_PARTITION_VERSION",
				ErrorMessage: "Invalid partition version: " + version,
			},
		}
	}

	pi.ColName = getCollectionName(version)
	pi.Instance.SetDB(pi.database)
	return pi.Instance.QueryWithOpt(context.TODO(), filter, opt)
}

func (pi *MonthlyPartitionInstance) Count(
	version string,
	query interface{},
) *common.Response {
	if !validPartitionVersion(version) {
		return &common.Response{
			Status: common.ResponseStatus.Error,
			Error: &common.ErrorResponse{
				ErrorCode:    "INVALID_PARTITION_VERSION",
				ErrorMessage: "Invalid partition version: " + version,
			},
		}
	}

	pi.ColName = getCollectionName(version)
	pi.Instance.SetDB(pi.database)
	return pi.Instance.Count(context.TODO(), query)
}

func (pi *MonthlyPartitionInstance) Upsert(
	version string,
	query interface{},
	updater interface{},
) *common.Response {
	if !validPartitionVersion(version) {
		return &common.Response{
			Status: common.ResponseStatus.Error,
			Error: &common.ErrorResponse{
				ErrorCode:    "INVALID_PARTITION_VERSION",
				ErrorMessage: "Invalid partition version: " + version,
			},
		}
	}

	pi.ColName = getCollectionName(version)
	pi.Instance.SetDB(pi.database)
	return pi.Instance.Upsert(context.TODO(), query, updater)
}

func (pi *MonthlyPartitionInstance) PrepareCol(version string) error {
	if !validPartitionVersion(version) {
		return fmt.Errorf("invalid partition version " + version)
	}

	pi.ColName = getCollectionName(version)
	pi.Instance.SetDB(pi.database)

	if !pi.existCol(version) {
		pi.database.CreateCollection(context.TODO(), pi.ColName)
	}

	err := pi.CreateIndex()
	if err != nil {
		return err
	}

	return nil
}

func (pi *MonthlyPartitionInstance) CreateIndex() error {

	return nil
}

func (pi *MonthlyPartitionInstance) existCol(version string) bool {
	existedColNames, err := pi.database.ListCollectionNames(context.TODO(), bson.M{})
	if err != nil {
		return false
	}

	colName := getCollectionName(version)
	for _, existedColName := range existedColNames {
		if existedColName == colName {
			return true
		}
	}

	return false
}

func getCollectionName(version string) string {
	return fmt.Sprintf("record_%s", version)
}

func validPartitionVersion(version string) bool {
	if len(version) == 0 {
		return false
	}

	sl := strings.Split(version, "_")
	if len(sl) != 2 {
		return false
	}

	_, err := strconv.Atoi(sl[0])
	if err != nil {
		return false
	}

	_, err = strconv.Atoi(sl[1])
	return err == nil
}
