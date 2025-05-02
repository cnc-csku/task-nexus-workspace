package handler

import "github.com/cnc-csku/task-nexus-workspace/internal/responses"

type HealthCheckHandler interface {
	HealthCheck() *responses.HealthCheckResponse
}

type healthCheckHandlerImpl struct{}

func NewHealthCheckHandler() HealthCheckHandler {
	return &healthCheckHandlerImpl{}
}

func (h *healthCheckHandlerImpl) HealthCheck() *responses.HealthCheckResponse {
	return &responses.HealthCheckResponse{
		Success: true,
	}
}
