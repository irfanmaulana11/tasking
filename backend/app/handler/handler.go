package handler

import "be-tasking/app/service"

type ApiHandler struct {
	healthCheckService service.HealthCheckServiceInterface
	authService        service.AuthServiceInterface
	taskService        service.TaskServiceInterface
}

func NewApiHandler(
	healthCheckService service.HealthCheckServiceInterface,
	authService service.AuthServiceInterface,
	taskService service.TaskServiceInterface,
) *ApiHandler {
	return &ApiHandler{
		healthCheckService: healthCheckService,
		authService:        authService,
		taskService:        taskService,
	}
}
