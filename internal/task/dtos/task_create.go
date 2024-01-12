package dtos

import (
	"myapp/config"
	"myapp/internal/task/entity"
	gStruct "myapp/pkg/globalStruct"
	"time"
)

type CreateTaskRequest struct {
	UserID      int    `json:"user_id" validate:"required"`
	WorkerID    int    `json:"worker_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

func NewCreateTask(data CreateTaskRequest, cfg config.Config) entity.Task {
	now := time.Now()
	return entity.Task{
		WorkerID:    data.WorkerID,
		Title:       data.Title,
		Description: data.Description,

		Tracing: gStruct.Tracing{
			CreatedAt: &now,
			CreatedBy: data.UserID,
		},
	}
}
