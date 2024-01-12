package dtos

import "myapp/internal/task/entity"

type ListTaskRequest struct {
	WorkerIDs    []int               `json:"worker_ids"`
	Statuses     []entity.TaskStatus `json:"statuses"`
	CreatedByIDs []int               `json:"created_by_ids"`
}
