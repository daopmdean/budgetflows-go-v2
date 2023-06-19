package biz

import (
	"context"
	"time"

	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/daopmdean/budgetflows-go-v2/model"
	"github.com/daopmdean/budgetflows-go-v2/utils"
	"github.com/daopmdean/summer/common"
	"github.com/gin-gonic/gin"
)

func Login(data *model.LoginRequest) *common.Response {
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

	queryRes := entity.AppUserDB.QueryOne(context.TODO(), &entity.AppUser{
		Username: data.Username,
	})
	if queryRes.Status != common.ResponseStatus.Success {
		return &common.Response{
			Status:  common.ResponseStatus.Invalid,
			Message: "Invalid username or password",
		}
	}

	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return &common.Response{
			Status: common.ResponseStatus.Error,
			Error: &common.ErrorResponse{
				ErrorCode:    "REGISTER_FAILED",
				ErrorMessage: err.Error(),
			},
		}
	}

	if validPass := utils.CheckPasswordHash(data.Password, hashedPassword); !validPass {
		return &common.Response{
			Status:  common.ResponseStatus.Invalid,
			Message: "Invalid username or password",
		}
	}

	token, err := GenerateToken(queryRes.Data.([]*entity.AppUser)[0], 24*time.Hour)
	if err != nil {
		return &common.Response{
			Status: common.ResponseStatus.Error,
			Error: &common.ErrorResponse{
				ErrorCode:    "REGISTER_FAILED",
				ErrorMessage: err.Error(),
			},
		}
	}

	return &common.Response{
		Status: common.ResponseStatus.Success,
		Data: []any{
			gin.H{
				"message": "Register success!",
				"user":    queryRes.Data.([]*entity.AppUser)[0],
				"token":   token,
			},
		},
	}
}