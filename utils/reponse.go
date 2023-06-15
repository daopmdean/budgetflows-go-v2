package utils

import "github.com/daopmdean/summer/common"

func Unauthorized() *common.Response {
	return &common.Response{
		Status: common.ResponseStatus.Unauthorized,
		Error: &common.ErrorResponse{
			ErrorCode:    "UNAUTHORIZED",
			ErrorMessage: "Unauthorized",
		},
	}
}

func Invalid(msg string) *common.Response {
	return &common.Response{
		Status: common.ResponseStatus.Invalid,
		Error: &common.ErrorResponse{
			ErrorCode:    "INVALID",
			ErrorMessage: msg,
		},
	}
}
