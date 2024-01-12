package repository

import (
	"context"
	"myapp/internal/worker/entity"
)

func (r *repository) GetWorkerByID(ctx context.Context, id uint) (result entity.Worker, err error) {
	var db = r.db.WithContext(ctx)

	condition := &entity.Worker{
		ID: id,
	}

	query := db.First(&result, condition)
	if query.Error != nil {
		r.log.Z().Err(query.Error).Msg("worker.GetWorkerByID")
		return
	}

	return
}
