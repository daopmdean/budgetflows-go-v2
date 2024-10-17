package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRole struct {
	ID          *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	CreatedTime *time.Time          `json:"createdTime,omitempty" bson:"created_time,omitempty"`
	UpdatedTime *time.Time          `json:"updatedTime,omitempty" bson:"updated_time,omitempty"`

	UserId *primitive.ObjectID `json:"userId,omitempty" bson:"user_id,omitempty"`
	Roles  []*UserRoleDetail   `json:"roles,omitempty" bson:"roles,omitempty"`
}

type UserRoleDetail struct {
	RoleId     *primitive.ObjectID `json:"roleId,omitempty" bson:"role_id,omitempty"`
	ValidUntil *time.Time          `json:"validUntil,omitempty" bson:"valid_until,omitempty"`
}
