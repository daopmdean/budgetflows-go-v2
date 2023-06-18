package biz

import (
	"strings"
)

func GetToken(authHeader string) string {
	headers := strings.Split(authHeader, " ")
	if len(headers) != 2 {
		return ""
	}

	if headers[0] != "Bearer" {
		return ""
	}

	return headers[1]
}
