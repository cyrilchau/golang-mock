package repository

import (
	"context"
	"myapp/internal/user/entity"
)

func (r *repository) GetUserByEmail(ctx context.Context, email string) (result entity.User, err error) {
	var db = r.db.WithContext(ctx)

	condition := &entity.User{
		Email: email,
	}

	query := db.Find(&result, condition)
	if query.Error != nil {
		r.log.Z().Err(query.Error).Msg("user.GetUserByEmail")
		return
	}

	return
}
