package router

import "github.com/labstack/echo/v4"

func (r *Router) RegisterAPIRouter(e *echo.Echo) {
	apiV1 := e.Group("/api/v1")

	apiV1.GET("/health", r.healthCheck.HealthCheck)
}
