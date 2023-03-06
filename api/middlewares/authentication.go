package middlewares

import (
	"marcelofelixsalgado/financial-period-api/api/responses"
	"marcelofelixsalgado/financial-period-api/api/responses/faults"
	"marcelofelixsalgado/financial-period-api/commons/logger"
	"marcelofelixsalgado/financial-period-api/pkg/infrastructure/auth"

	"github.com/labstack/echo/v4"
)

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := auth.ValidateToken(c.Request()); err != nil {
			logger.GetLogger().Infof("Token validation error: %v", err)
			responseMessage := responses.NewResponseMessage().AddMessageByErrorCode(faults.NotAuthorized)
			return c.JSON(responseMessage.HttpStatusCode, responseMessage)
		}
		return next(c)
	}
}