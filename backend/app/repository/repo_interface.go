package repository

import (
	"be-tasking/app/model"
	"context"
)

type MySQLRepoInterface interface {
	// user
	CreateUser(ctx context.Context, req model.User) error
	GetUserByUserName(ctx context.Context, username string) (model.User, error)

	// task
	CreateTask(ctx context.Context, req model.Task) error
	GetTaskByID(ctx context.Context, id string) (model.Task, error)
	UpdateTask(ctx context.Context, taskId string, data map[string]interface{}) error
	GetTaskList(ctx context.Context, filter model.TableFilter) ([]model.Task, int, error)

	// task hostory
	CreateTaskHistory(ctx context.Context, req model.TaskHistory) error
	GetTaskHistory(ctx context.Context, taskId string) ([]model.TaskHistory, error)
}
