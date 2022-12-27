package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type IHealthHandler interface {
	Health(ctx echo.Context) error
}

type message struct {
	Status string `json:"status"`
}

type HealthHandler struct {
}

func NewHealthHandler() IHealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(ctx echo.Context) error {

	successMessage := message{
		Status: "Ok",
	}

	return ctx.JSON(http.StatusOK, successMessage)
}
