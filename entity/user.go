package entity

import (
	"time"

	"github.com/daopmdean/summer/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var UserDB = &mongodb.Instance{
	ColName:     "user",
	TemplateObj: &User{},
}

type User struct {
	ID          *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	CreatedTime *time.Time          `json:"createdTime,omitempty" bson:"created_time,omitempty"`
	UpdatedTime *time.Time          `json:"updatedTime,omitempty" bson:"updated_time,omitempty"`

	UserId   int64  `json:"userId,omitempty" bson:"user_id,omitempty"`
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Phone    string `json:"phone,omitempty" bson:"phone,omitempty"`

	Password     string `json:"password,omitempty" bson:"password,omitempty"`
	PasswordSalt string `json:"passwordSalt,omitempty" bson:"password_salt,omitempty"`

	Name    string     `json:"name,omitempty" bson:"name,omitempty"`
	Address string     `json:"address,omitempty" bson:"address,omitempty"`
	Avatar  string     `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Dob     *time.Time `json:"dob,omitempty" bson:"dob,omitempty"`
}
