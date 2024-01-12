package usecase

import (
	"context"
	"myapp/config"
	"myapp/internal/worker/dtos"
	"myapp/internal/worker/entity"
	"myapp/internal/worker/repository"
	"myapp/pkg/otel/zerolog"
)

type (
	Usecase interface {
		CreateOneWorker(ctx context.Context, payload dtos.CreateWorkerRequest) (result entity.Worker, httpCode int, err error)
		DetailWorker(ctx context.Context, id uint) (detail entity.Worker, httpCode int, err error)
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
