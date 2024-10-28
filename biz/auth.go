package biz

import (
	"github.com/daopmdean/budgetflows-go-v2/entity"
	"github.com/daopmdean/summer/common"
)

func Require(role string, userRole *entity.UserRole) *common.Response {
	if userRole == nil {
		return common.BuildUnauthorizedRes("UNAUTHORIZED", "missing user role")
	}

	for _, roleDetail := range userRole.Roles {
		if roleDetail.Name == role {
			return &common.Response{
				Status:  common.ResponseStatus.Success,
				Message: "Success",
			}
		}
	}

	return common.BuildUnauthorizedRes("UNAUTHORIZED", "no valid role")
}
