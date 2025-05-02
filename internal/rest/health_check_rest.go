package rest

import (
	"net/http"

	"github.com/cnc-csku/task-nexus-workspace/internal/handler"
	"github.com/labstack/echo/v4"
)

type HealthCheckRest interface {
	HealthCheck(c echo.Context) error
}

type healthCheckRestImpl struct {
	healthCheckHandler handler.HealthCheckHandler
}

func NewHealthCheckRest(healthCheckHandler handler.HealthCheckHandler) HealthCheckRest {
	return &healthCheckRestImpl{
		healthCheckHandler: healthCheckHandler,
	}
}

func (h *healthCheckRestImpl) HealthCheck(c echo.Context) error {
	res := h.healthCheckHandler.HealthCheck()

	return c.JSON(http.StatusOK, res)
}
