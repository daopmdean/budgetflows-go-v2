package biz

import (
	"context"
	"strings"
	"time"

	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/daopmdean/budgetflows-go-v2/model"
	"github.com/daopmdean/summer/auth"
	"github.com/daopmdean/summer/common"
	"github.com/gin-gonic/gin"
)

func Register(data *model.RegisterRequest) *common.Response {
	if data.Username == "" {
		return &common.Response{
			Status:  common.ResponseStatus.Invalid,
			Message: "Username is required",
		}
	}

	if data.Password == "" {
		return &common.Response{
			Status:  common.ResponseStatus.Invalid,
			Message: "Password is required",
		}
	}

	data.Username = strings.ToLower(data.Username)

	queryRes := entity.AppUserDB.QueryOne(context.TODO(), &entity.AppUser{
		Username: data.Username,
	})
	if queryRes.Status != common.ResponseStatus.NotFound {
		return &common.Response{
			Status:  common.ResponseStatus.Invalid,
			Message: "Username is already taken",
		}
	}

	hashedPassword, err := auth.GetHashed(data.Password)
	if err != nil {
		return common.BuildErrorRes("REGISTER_FAILED", err.Error())
	}

	registerRes := entity.AppUserDB.Create(context.TODO(), &entity.AppUser{
		Username: data.Username,
		UserId:   entity.GenUserId(),
		Phone:    data.Phone,
		Email:    data.Email,
		Password: hashedPassword,

		Name:    data.Name,
		Address: data.Address,
		Avatar:  data.Avatar,
		Dob:     data.Dob,
	})
	if registerRes.Status != common.ResponseStatus.Success {
		return registerRes
	}

	token, err := GenerateToken(registerRes.Data.([]*entity.AppUser)[0], 24*time.Hour)
	if err != nil {
		return common.BuildErrorRes("REGISTER_FAILED", err.Error())
	}

	return &common.Response{
		Status: common.ResponseStatus.Success,
		Data: []any{
			gin.H{
				"message": "Register success!",
				"user":    registerRes.Data.([]*entity.AppUser)[0],
				"token":   token,
			},
		},
	}
}
