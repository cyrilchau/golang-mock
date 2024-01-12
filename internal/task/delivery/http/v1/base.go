package delivery

import (
	"myapp/pkg/otel/zerolog"

	taskUsecase "myapp/internal/task/usecase"

	"github.com/labstack/echo/v4"
)

type (
	Handlers interface {
		Create(c echo.Context) error
		Detail(c echo.Context) error
		List(c echo.Context) error
	}

	handlers struct {
		uc  taskUsecase.Usecase
		log *zerolog.Logger
	}

	HandlersDeps struct {
		TaskUsecaseI taskUsecase.Usecase
	}
)

const (
	BooleanTextTrue  = "true"
	BooleanTextFalse = "false"
)

func NewHandlers(deps HandlersDeps, log *zerolog.Logger) Handlers {
	return &handlers{
		uc:  deps.TaskUsecaseI,
		log: log,
	}
}
