package service

import (
	"be-tasking/app/model"
	"be-tasking/app/service/dto"

	"github.com/gin-gonic/gin"
)

type HealthCheckServiceInterface interface {
	Check() dto.HealthCheck
}

type AuthServiceInterface interface {
	Register(c *gin.Context, req dto.Register) (int, error)
	Login(c *gin.Context, login dto.Login) (int, *dto.LoginResponse, error)
}

type TaskServiceInterface interface {
	CreateTask(c *gin.Context, req dto.Task) (int, *dto.Task, error)
	UpdateTask(c *gin.Context, req dto.Task) (int, *dto.Task, error)
	UpdateTaskStatus(c *gin.Context, req dto.TaskProgress) (int, *dto.TaskProgress, error)
	GetTaskList(c *gin.Context, filter model.TableFilter) (int, []dto.Task, interface{}, error)
}
