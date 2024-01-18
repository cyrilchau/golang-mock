package delivery

import (
	"myapp/pkg/otel/zerolog"

	userUsecase "myapp/internal/user/usecase"

	"github.com/labstack/echo/v4"
)

type (
	Services interface {
		Login(c echo.Context) error
		Create(c echo.Context) error
		Detail(c echo.Context) error
	}

	service struct {
		uc  userUsecase.Usecase
		log *zerolog.Logger
	}

	ServiceDeps struct {
		UserUsecaseI userUsecase.Usecase
	}
)

func NewService(deps ServiceDeps, log *zerolog.Logger) Services {
	return &service{
		uc:  deps.UserUsecaseI,
		log: log,
	}
}
