package router

import "github.com/cnc-csku/task-nexus-workspace/internal/rest"

type Router struct {
	healthCheck rest.HealthCheckRest
}

func NewRouter(
	healthCheck rest.HealthCheckRest,
) *Router {
	return &Router{
		healthCheck: healthCheck,
	}
}
