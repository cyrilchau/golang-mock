package delivery

import (
	"myapp/pkg/otel/zerolog"

	workerUsecase "myapp/internal/worker/usecase"

	"github.com/labstack/echo/v4"
)

type (
	Handlers interface {
		Create(c echo.Context) error
		Detail(c echo.Context) error
	}

	handlers struct {
		uc  workerUsecase.Usecase
		log *zerolog.Logger
	}

	HandlersDeps struct {
		WorkerUsecaseI workerUsecase.Usecase
	}
)

const (
	BooleanTextTrue  = "true"
	BooleanTextFalse = "false"
)

func NewHandlers(deps HandlersDeps, log *zerolog.Logger) Handlers {
	return &handlers{
		uc:  deps.WorkerUsecaseI,
		log: log,
	}
}
