package delivery

import (
	"myapp/pkg/otel/zerolog"

	taskUsecase "myapp/internal/task/usecase"

	"github.com/labstack/echo/v4"
)

type (
	Services interface {
		Create(c echo.Context) error
		Detail(c echo.Context) error
		List(c echo.Context) error
	}

	service struct {
		uc  taskUsecase.Usecase
		log *zerolog.Logger
	}

	ServiceDeps struct {
		TaskUsecaseI taskUsecase.Usecase
	}
)

func NewService(deps ServiceDeps, log *zerolog.Logger) Services {
	return &service{
		uc:  deps.TaskUsecaseI,
		log: log,
	}
}
