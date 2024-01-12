package repository

import (
	"context"
	"myapp/internal/task/entity"
)

func (r *repository) GetTaskByID(ctx context.Context, id uint) (result entity.Task, err error) {
	var db = r.db.WithContext(ctx)

	condition := &entity.Task{
		ID: id,
	}

	query := db.Find(&result, condition)
	if query.Error != nil {
		r.log.Z().Err(query.Error).Msg("user.GetTaskByID")
		return
	}

	return
}
