package repository

import (
	"context"
	"myapp/internal/task/dtos"
	"myapp/internal/task/entity"
	"myapp/pkg/otel/zerolog"

	"gorm.io/gorm"
)

type (
	Repository interface {
		CreateNewTask(context.Context, entity.Task) (entity.Task, error)
		GetAllTask(context.Context, dtos.ListTaskRequest) ([]entity.Task, error)
		GetTaskByID(context.Context, uint) (entity.Task, error)
	}

	repository struct {
		db  *gorm.DB
		log *zerolog.Logger
	}
)

func NewRepository(c *gorm.DB, l *zerolog.Logger) Repository {
	return &repository{log: l, db: c}
}
