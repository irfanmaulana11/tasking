package dbrepository

import (
	"be-tasking/app/model"
	"context"
)

func (r *mySQLRepository) CreateTaskHistory(ctx context.Context, req model.TaskHistory) error {
	var err error

	tx := r.db.Begin()
	if err = tx.Table("task_history").Create(&req).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (r *mySQLRepository) GetTaskHistory(ctx context.Context, taskId string) ([]model.TaskHistory, error) {
	var (
		hst []model.TaskHistory
		err error
	)

	tx := r.db.Begin()
	if err = tx.Table("task_history").Find(&hst, "task_id = ?", taskId).Error; err != nil {
		tx.Rollback()
		return hst, err
	}
	tx.Commit()

	return hst, nil
}
