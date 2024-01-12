package delivery

import (
	"myapp/pkg/otel/zerolog"

	userUsecase "myapp/internal/user/usecase"

	"github.com/labstack/echo/v4"
)

type (
	Handlers interface {
		Login(c echo.Context) error
		Create(c echo.Context) error
		Detail(c echo.Context) error
	}

	handlers struct {
		uc  userUsecase.Usecase
		log *zerolog.Logger
	}

	HandlersDeps struct {
		UserUsecaseI userUsecase.Usecase
	}
)

const (
	BooleanTextTrue  = "true"
	BooleanTextFalse = "false"
)

func NewHandlers(deps HandlersDeps, log *zerolog.Logger) Handlers {
	return &handlers{
		uc:  deps.UserUsecaseI,
		log: log,
	}
}
