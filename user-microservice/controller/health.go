package controller

import (
	"net/http"
	"t/container"

	"github.com/labstack/echo/v4"
)

type HealthController struct {
	container container.Container
}

func NewHealthController(container container.Container) *HealthController {
	return &HealthController{container: container}
}

func (controller *HealthController) GetHealth(c echo.Context) error {

	return c.JSON(http.StatusOK, "healthy")
}
