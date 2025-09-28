package service

import (
	"be-tasking/app/repository"
)

type healthCheckService struct{}

type authService struct {
	MySQL repository.MySQLRepoInterface
}

type taskService struct {
	MySQL repository.MySQLRepoInterface
}

type services struct {
	HealthCheck HealthCheckServiceInterface
	Auth        AuthServiceInterface
	Task        TaskServiceInterface
}

func NewHealthCheckService() HealthCheckServiceInterface {
	return &healthCheckService{}
}

func NewAuthService(mysql repository.MySQLRepoInterface) AuthServiceInterface {
	return &authService{
		MySQL: mysql,
	}
}

func NewTaskService(mysql repository.MySQLRepoInterface) TaskServiceInterface {
	return &taskService{
		MySQL: mysql,
	}
}

func NewServices(repo repository.MySQLRepoInterface) *services {
	return &services{
		HealthCheck: NewHealthCheckService(),
		Auth:        NewAuthService(repo),
		Task:        NewTaskService(repo),
	}
}
