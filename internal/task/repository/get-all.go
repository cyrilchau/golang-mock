package repository

import (
	"context"
	"myapp/internal/task/dtos"
	"myapp/internal/task/entity"
)

func (r *repository) GetAllTask(ctx context.Context, req dtos.ListTaskRequest) (result []entity.Task, err error) {
	var db = r.db.WithContext(ctx)

    conditions := make(map[string]interface{})

	if len(req.WorkerIDs) > 0 {
		conditions["worker_id"] = req.WorkerIDs
	}
	if len(req.Statuses) > 0 {
		conditions["status"] = req.Statuses
	}
	if len(req.CreatedByIDs) > 0 {
		conditions["created_by"] = req.CreatedByIDs
	}

	query := db.Where(conditions).Find(&result)
	if query.Error != nil {
		r.log.Z().Err(query.Error).Msg("user.GetTaskByID")
		return
	}

	return
}
