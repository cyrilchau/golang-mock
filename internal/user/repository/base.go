package repository

import (
	"context"
	"myapp/internal/user/entity"
	"myapp/pkg/otel/zerolog"

	"gorm.io/gorm"
)

type (
	Repository interface {
		GetUserByEmail(context.Context, string) (entity.User, error)
		GetUserByID(context.Context, int) (entity.User, error)
		CreateOneUser(context.Context, entity.User) (entity.User, error)
	}

	repository struct {
		db  *gorm.DB
		log *zerolog.Logger
	}
)

func NewRepository(c *gorm.DB, l *zerolog.Logger) Repository {
	return &repository{log: l, db: c}
}
