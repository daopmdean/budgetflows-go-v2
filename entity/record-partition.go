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
	partition map[string]*mongodb.Instance // key: collection name, value: mongodb instance
	database  *mongo.Database
}

var RecordDBPartition = &MonthlyPartitionInstance{}

func (pi *MonthlyPartitionInstance) GetPartitions() map[string]*mongodb.Instance {
	return pi.partition
}

func (pi *MonthlyPartitionInstance) SetDBPartition(
	database *mongo.Database,
) *MonthlyPartitionInstance {
	pi.database = database
	pi.partition = make(map[string]*mongodb.Instance)
	return pi
}

func (pi *MonthlyPartitionInstance) Create(
	version string,
	ele interface{},
) *common.Response {
	if !validPartitionVersion(version) {
		return common.BuildErrorRes(
			"INVALID_PARTITION_VERSION",
			"Invalid partition version: "+version,
		)
	}

	ins := pi.getPartition(version)
	return ins.Create(context.TODO(), ele)
}

func (pi *MonthlyPartitionInstance) Query(
	version string,
	filter interface{},
	skip, limit int64,
	sortFields *primitive.M,
) *common.Response {
	if !validPartitionVersion(version) {
		return common.BuildErrorRes(
			"INVALID_PARTITION_VERSION",
			"Invalid partition version: "+version,
		)
	}

	ins := pi.getPartition(version)
	return ins.Query(context.TODO(), filter, skip, limit, sortFields)
}

func (pi *MonthlyPartitionInstance) QueryOne(
	version string,
	filter interface{},
) *common.Response {
	if !validPartitionVersion(version) {
		return common.BuildErrorRes(
			"INVALID_PARTITION_VERSION",
			"Invalid partition version: "+version,
		)
	}

	ins := pi.getPartition(version)
	return ins.QueryOne(context.TODO(), filter)
}

func (pi *MonthlyPartitionInstance) QueryWithOpt(
	version string,
	filter interface{},
	opt *options.FindOptions,
) *common.Response {
	if !validPartitionVersion(version) {
		return common.BuildErrorRes(
			"INVALID_PARTITION_VERSION",
			"Invalid partition version: "+version,
		)
	}

	ins := pi.getPartition(version)
	return ins.QueryWithOpt(context.TODO(), filter, opt)
}

func (pi *MonthlyPartitionInstance) Count(
	version string,
	query interface{},
) *common.Response {
	if !validPartitionVersion(version) {
		return common.BuildErrorRes(
			"INVALID_PARTITION_VERSION",
			"Invalid partition version: "+version,
		)
	}

	ins := pi.getPartition(version)
	return ins.Count(context.TODO(), query)
}

func (pi *MonthlyPartitionInstance) Upsert(
	version string,
	query interface{},
	updater interface{},
) *common.Response {
	if !validPartitionVersion(version) {
		return common.BuildErrorRes(
			"INVALID_PARTITION_VERSION",
			"Invalid partition version: "+version,
		)
	}

	ins := pi.getPartition(version)
	return ins.Upsert(context.TODO(), query, updater)
}

func (pi *MonthlyPartitionInstance) Delete(
	version string,
	filter interface{},
) *common.Response {
	if !validPartitionVersion(version) {
		return common.BuildErrorRes(
			"INVALID_PARTITION_VERSION",
			"Invalid partition version: "+version,
		)
	}

	ins := pi.getPartition(version)
	return ins.Delete(context.TODO(), filter)
}

func (pi *MonthlyPartitionInstance) PrepareCol(version string) error {
	if !validPartitionVersion(version) {
		return fmt.Errorf("invalid partition version " + version)
	}

	if !pi.existCol(version) {
		pi.database.CreateCollection(context.TODO(), getCollectionName(version))
	}

	err := pi.CreateIndex(version)
	if err != nil {
		return err
	}

	return nil
}

func (pi *MonthlyPartitionInstance) CreateIndex(version string) error {
	pi.database.
		Client().
		Database(pi.database.Name()).
		Collection(getCollectionName(version)).
		Indexes().
		CreateOne(context.TODO(), mongo.IndexModel{
			Keys: bson.D{
				{Key: "user_id", Value: 1},
				{Key: "tags", Value: 1},
			},
		})

	return nil
}

func (pi *MonthlyPartitionInstance) getPartition(version string) *mongodb.Instance {
	colName := getCollectionName(version)
	if ins, ok := pi.partition[colName]; ok {
		return ins
	}

	pi.partition[colName] = &mongodb.Instance{
		ColName:     colName,
		TemplateObj: new(Record),
	}
	pi.partition[colName].SetDB(pi.database)

	return pi.partition[colName]
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
