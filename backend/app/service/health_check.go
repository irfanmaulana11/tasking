package service

import (
	"be-tasking/app/service/dto"
	"net/http"
	"time"
)

func (s *healthCheckService) Check() dto.HealthCheck {
	healthCheck := dto.HealthCheck{
		Message: http.StatusText(http.StatusOK),
		Time:    time.Now(),
	}
	return healthCheck
}
