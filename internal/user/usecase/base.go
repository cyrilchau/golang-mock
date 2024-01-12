package usecase

import (
	"context"
	"myapp/config"
	"myapp/internal/user/dtos"
	"myapp/internal/user/entity"
	"myapp/internal/user/repository"
	"myapp/pkg/otel/zerolog"

	"github.com/golang-jwt/jwt/v5"
)

type (
	Usecase interface {
		Login(ctx context.Context, request dtos.UserLoginRequest) (response dtos.UserLoginResponse, httpCode int, err error)
		Create(ctx context.Context, payload dtos.CreateUserRequest) (result entity.User, httpCode int, err error)
		Detail(ctx context.Context, id int) (detail dtos.UserDetailResponse, httpCode int, err error)

		IntrospectToken(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error)
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
