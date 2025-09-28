package dbrepository

import (
	"be-tasking/app/model"
	"context"
)

func (r *mySQLRepository) CreateUser(ctx context.Context, req model.User) error {
	var err error

	tx := r.db.Begin()
	if err = tx.Create(&req).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (r *mySQLRepository) GetUserByUserName(ctx context.Context, username string) (model.User, error) {
	var (
		user model.User
		err  error
	)

	tx := r.db.Begin()
	if err = tx.Find(&user, "username = ?", username).Error; err != nil {
		tx.Rollback()
		return user, err
	}
	tx.Commit()

	return user, nil
}
