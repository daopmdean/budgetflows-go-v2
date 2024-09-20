package rest

import (
	"net/http"

	"github.com/daopmdean/budgetflows-go-v2/conf"
	"github.com/daopmdean/summer/auth"
	"github.com/daopmdean/summer/common"
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, res *common.Response) {
	if res.Status == common.ResponseStatus.Success {
		c.JSON(http.StatusOK, res)
	} else if res.Status == common.ResponseStatus.NotFound {
		c.JSON(http.StatusNotFound, res)
	} else if res.Status == common.ResponseStatus.Invalid {
		c.JSON(http.StatusBadRequest, res)
	} else if res.Status == common.ResponseStatus.Unauthorized {
		c.JSON(http.StatusUnauthorized, res)
	} else if res.Status == common.ResponseStatus.Error {
		c.JSON(http.StatusInternalServerError, res)
	}
}

func getClaims(c *gin.Context) (*auth.SummerClaim, error) {
	bearer := c.Request.Header.Get("Authorization")
	token, err := auth.ExtractTokenFromHeader(bearer)
	if err != nil {
		return nil, err
	}

	return auth.ParseToken(token, conf.AppConfig.SignedKey)
}
