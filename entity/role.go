package entity

import (
	"time"

	"github.com/daopmdean/summer/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var RoleDB = &mongodb.Instance{
	ColName:     "role",
	TemplateObj: &Role{},
}

type Role struct {
	ID          *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	CreatedTime *time.Time          `json:"createdTime,omitempty" bson:"created_time,omitempty"`
	UpdatedTime *time.Time          `json:"updatedTime,omitempty" bson:"updated_time,omitempty"`

	RoleId      int64  `json:"roleId,omitempty" bson:"role_id,omitempty"`
	RoleCode    string `json:"roleCode,omitempty" bson:"role_code,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
}
