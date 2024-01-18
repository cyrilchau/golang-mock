package delivery

import (
	"myapp/pkg/otel/zerolog"

	workerUsecase "myapp/internal/worker/usecase"

	"github.com/labstack/echo/v4"
)

type (
	Services interface {
		Create(c echo.Context) error
		Detail(c echo.Context) error
	}

	service struct {
		uc  workerUsecase.Usecase
		log *zerolog.Logger
	}

	ServiceDeps struct {
		WorkerUsecaseI workerUsecase.Usecase
	}
)

func NewService(deps ServiceDeps, log *zerolog.Logger) Services {
	return &service{
		uc:  deps.WorkerUsecaseI,
		log: log,
	}
}
