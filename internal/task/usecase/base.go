package usecase

import (
	"context"
	"myapp/config"
	"myapp/internal/task/dtos"
	"myapp/internal/task/entity"
	"myapp/internal/task/repository"
	"myapp/pkg/otel/zerolog"
)

type (
	Usecase interface {
		CreateNewTask(ctx context.Context, payload dtos.CreateTaskRequest) (result entity.Task, httpCode int, err error)
		GetTaskDetail(ctx context.Context, payload dtos.DetailTaskRequest) (result entity.Task, httpCode int, err error)
		GetListTask(ctx context.Context, payload dtos.ListTaskRequest) (result []entity.Task, httpCode int, err error)
	}

	usecase struct {
		repo repository.Repository
		log  *zerolog.Logger
		cfg  config.Config
	}
)

func NewUseCase(repo repository.Repository, log *zerolog.Logger, cfg config.Config) Usecase {
	return &usecase{repo, log, cfg}
}
