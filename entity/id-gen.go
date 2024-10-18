package entity

import (
	"context"

	"github.com/daopmdean/budgetflows-go-v2/utils"
	"github.com/daopmdean/summer/mongodb"
)

var IdGenDB = &mongodb.Instance{
	ColName:     "_id_gen",
	TemplateObj: &IdGen{},
}

type IdGen struct {
	ID    string `json:"id,omitempty" bson:"_id,omitempty"`
	Value int64  `json:"value,omitempty" bson:"value,omitempty"`
}

func GenUserId() int64 {
	increResult := IdGenDB.IncreValue(context.TODO(), IdGen{
		ID: "APP_USER",
	}, "value", 1)
	val := increResult.Data.([]*IdGen)[0]
	return val.Value
}

func GenRoleCode() (int64, string) {
	increResult := IdGenDB.IncreValue(context.TODO(), IdGen{
		ID: "ROLE",
	}, "value", 1)
	val := increResult.Data.([]*IdGen)[0]
	return val.Value, utils.GetCode(val.Value, 3)
}
