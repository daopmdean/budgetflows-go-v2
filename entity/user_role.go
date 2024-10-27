package entity

import (
	"time"

	"github.com/daopmdean/summer/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var UserRoleDB = &mongodb.Instance{
	ColName:     "user_role",
	TemplateObj: &UserRole{},
}

type UserRole struct {
	ID          *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	CreatedTime *time.Time          `json:"createdTime,omitempty" bson:"created_time,omitempty"`
	UpdatedTime *time.Time          `json:"updatedTime,omitempty" bson:"updated_time,omitempty"`

	UserId int64             `json:"userId,omitempty" bson:"user_id,omitempty"`
	Roles  []*UserRoleDetail `json:"roles,omitempty" bson:"roles,omitempty"`
}

type UserRoleDetail struct {
	RoleId     int64      `json:"roleId,omitempty" bson:"role_id,omitempty"`
	RoleCode   string     `json:"roleCode,omitempty" bson:"role_code,omitempty"`
	Name       string     `json:"name,omitempty" bson:"name,omitempty"`
	ValidUntil *time.Time `json:"validUntil,omitempty" bson:"valid_until,omitempty"`
}
