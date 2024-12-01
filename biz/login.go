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

func Login(req *model.LoginRequest) *common.Response {
	if req.Username == "" {
		return &common.Response{
			Status:  common.ResponseStatus.Invalid,
			Message: "Username is required",
		}
	}

	if req.Password == "" {
		return &common.Response{
			Status:  common.ResponseStatus.Invalid,
			Message: "Password is required",
		}
	}

	queryRes := entity.UserDB.QueryOne(context.TODO(), &entity.User{
		Username: strings.ToLower(req.Username),
	})
	if queryRes.Status != common.ResponseStatus.Success {
		return &common.Response{
			Status:  common.ResponseStatus.Invalid,
			Message: "Invalid username or password",
		}
	}

	user := queryRes.Data.([]*entity.User)[0]
	if err := auth.CheckPasswordHash(req.Password+user.PasswordSalt, user.Password); err != nil {
		return &common.Response{
			Status:  common.ResponseStatus.Invalid,
			Message: "Invalid username or password",
		}
	}

	token, err := GenerateToken(queryRes.Data.([]*entity.User)[0], 300*24*time.Hour)
	if err != nil {
		return common.BuildErrorRes("LOGIN_FAILED", err.Error())
	}

	return &common.Response{
		Status: common.ResponseStatus.Success,
		Data: []any{
			gin.H{
				"message": "Login success!",
				"user":    queryRes.Data.([]*entity.User)[0],
				"token":   token,
			},
		},
	}
}
