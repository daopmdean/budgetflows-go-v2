package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	ID          *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	CreatedTime *time.Time          `json:"createdTime,omitempty" bson:"created_time,omitempty"`
	UpdatedTime *time.Time          `json:"updatedTime,omitempty" bson:"updated_time,omitempty"`

	Name string `json:"name,omitempty" bson:"name,omitempty"`
}
