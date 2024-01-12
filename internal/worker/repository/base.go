package repository

import (
	"context"
	"myapp/internal/worker/entity"
	"myapp/pkg/otel/zerolog"

	"gorm.io/gorm"
)

type (
	Repository interface {
		CreateOneWorker(context.Context, entity.Worker) (entity.Worker, error)
		GetWorkerByID(context.Context, uint) (entity.Worker, error)
	}

	repository struct {
		db  *gorm.DB
		log *zerolog.Logger
	}
)

func NewRepository(c *gorm.DB, l *zerolog.Logger) Repository {
	return &repository{log: l, db: c}
}
