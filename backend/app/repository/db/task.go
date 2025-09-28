package dbrepository

import (
	"be-tasking/app/model"
	"be-tasking/constanta"
	"context"

	"github.com/asaskevich/govalidator"
)

func (r *mySQLRepository) CreateTask(ctx context.Context, req model.Task) error {
	var err error

	tx := r.db.Begin()
	if err = tx.Create(&req).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (r *mySQLRepository) GetTaskByID(ctx context.Context, id string) (model.Task, error) {
	var (
		task model.Task
		err  error
	)

	tx := r.db.Begin()
	if err = tx.Find(&task, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return task, err
	}
	tx.Commit()

	return task, nil
}

func (r *mySQLRepository) UpdateTask(ctx context.Context, taskId string, data map[string]interface{}) error {
	var err error

	tx := r.db.Begin()
	err = tx.Debug().Table("tasks").Where("id = ?", taskId).Updates(data).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (r *mySQLRepository) GetTaskList(ctx context.Context, filter model.TableFilter) ([]model.Task, int, error) {
	var (
		task []model.Task
	)

	tx := r.db.Begin().Table("tasks")
	if !govalidator.IsNull(filter.Search) {
		search := "%" + filter.Search + "%"
		tx = tx.Where("title LIKE ? OR description LIKE ?", search, search)
	}

	if filter.Role == constanta.RoleTypeManager {
		tx = tx.Where("status not in ('Submitted','Revision','Updated')")
	}

	// Fetch the total count
	var total int
	if err := tx.Model(&model.Task{}).Count(&total).Error; err != nil {
		tx.Rollback()
		return task, 0, err
	}

	offset := filter.Limit * (filter.Page - 1)

	// Fetch the filtered data
	if err := tx.Order("created_at DESC").Limit(filter.Limit).Offset(offset).Find(&task).Error; err != nil {
		tx.Rollback()
		return task, 0, err
	}
	tx.Commit()

	return task, total, nil
}
