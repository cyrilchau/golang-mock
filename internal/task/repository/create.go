package repository

import (
	"context"
	"myapp/internal/task/entity"
)

func (r *repository) CreateNewTask(ctx context.Context, task entity.Task) (result entity.Task, err error) {
	var db = r.db.WithContext(ctx)

	query := db.Create(&task)
	if query.Error != nil {
		err = query.Error
		r.log.Z().Err(query.Error).Msg("db.Create")
		return result, err
	}

	// Retrieve the created record
	err = query.Scan(&result).Error
	if err != nil {
		r.log.Z().Err(err).Msg("Error scanning created record")
		return result, err
	}

	return result, nil
}
