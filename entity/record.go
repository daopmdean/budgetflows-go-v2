package entity

import (
	"time"

	"github.com/daopmdean/summer/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var RecordDB = &mongodb.Instance{
	ColName:     "record",
	TemplateObj: &Record{},
}

type Record struct {
	ID          *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	CreatedTime *time.Time          `json:"createdTime,omitempty" bson:"created_time,omitempty"`
	UpdatedTime *time.Time          `json:"updatedTime,omitempty" bson:"updated_time,omitempty"`

	UserId     int64      `json:"userId,omitempty" bson:"user_id,omitempty"`
	Amount     float64    `json:"amount,omitempty" bson:"amount,omitempty"`
	Note       string     `json:"note,omitempty" bson:"note,omitempty"`
	Tags       []string   `json:"tags,omitempty" bson:"tags,omitempty"`
	RecordTime *time.Time `json:"recordTime,omitempty" bson:"record_time,omitempty"`

	TimeKey string `json:"timeKey,omitempty" bson:"-"`
}
