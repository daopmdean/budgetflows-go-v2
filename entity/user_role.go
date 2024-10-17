package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRole struct {
	ID          *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	CreatedTime *time.Time          `json:"createdTime,omitempty" bson:"created_time,omitempty"`
	UpdatedTime *time.Time          `json:"updatedTime,omitempty" bson:"updated_time,omitempty"`

	UserId  *primitive.ObjectID   `json:"userId,omitempty" bson:"user_id,omitempty"`
	RoleIds []*primitive.ObjectID `json:"roleIds,omitempty" bson:"role_ids,omitempty"`
}
