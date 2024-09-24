package entity

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/daopmdean/summer/common"
	"github.com/daopmdean/summer/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PartitionInstance struct {
	partitions    map[string]*mongodb.Instance // key: collection name, value: mongodb instance
	baseColName   string
	partitionType PartitionType
	database      *mongo.Database
}

type PartitionType string

var PartitionTypeValue = struct {
	Daily   PartitionType
	Monthly PartitionType
	Yearly  PartitionType
}{
	Daily:   "daily",
	Monthly: "monthly",
	Yearly:  "yearly",
}

func (pi *PartitionInstance) GetPartitions() map[string]*mongodb.Instance {
	return pi.partitions
}

func (pi *PartitionInstance) GetPartitionType() PartitionType {
	return pi.partitionType
}

func (pi *PartitionInstance) GetDatabase() *mongo.Database {
	return pi.database
}

func (pi *PartitionInstance) QueryWithOpt(
	t time.Time,
	filter interface{},
	opt *options.FindOptions,
) *common.Response {
	ins, err := pi.getPartition(t)
	if err != nil {
		return common.BuildErrorRes("GET_PARTITION_FAILED", err.Error())
	}

	return ins.QueryWithOpt(context.TODO(), filter, opt)
}

func (pi *PartitionInstance) getPartition(t time.Time) (*mongodb.Instance, error) {
	colName, err := pi.getCollectionName(t)
	if err != nil {
		return nil, err
	}

	if ins, ok := pi.partitions[colName]; ok {
		return ins, nil
	}

	pi.partitions[colName] = &mongodb.Instance{
		ColName:     colName,
		TemplateObj: new(Record),
	}
	pi.partitions[colName].SetDB(pi.database)

	return pi.partitions[colName], nil
}

func (pi *PartitionInstance) getCollectionName(t time.Time) (string, error) {
	version, err := getVersion(pi.partitionType, t)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s_%s", pi.baseColName, version), nil
}

func getVersion(partitionType PartitionType, t time.Time) (string, error) {
	switch partitionType {
	case PartitionTypeValue.Daily:
		return t.Format("2006_01_02"), nil
	case PartitionTypeValue.Monthly:
		return t.Format("2006_01"), nil
	case PartitionTypeValue.Yearly:
		return t.Format("2006"), nil
	}

	return "", errors.New("invalid partition type")
}
