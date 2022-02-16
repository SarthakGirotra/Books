package controller

import (
	"net/http"
	"t/container"
	"t/models"

	"github.com/labstack/echo/v4"
)

type HealthController struct {
	container container.Container
}

func NewHealthController(container container.Container) *HealthController {
	return &HealthController{container: container}
}

// GetHealth - check server health
func (controller *HealthController) GetHealth(c echo.Context) error {
	msg := &models.Response{Message: "healthy"}
	return c.JSON(http.StatusOK, msg)
}
