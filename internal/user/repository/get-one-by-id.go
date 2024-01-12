package repository

import (
	"context"
	"myapp/internal/user/entity"
)

func (r *repository) GetUserByID(ctx context.Context, id int) (result entity.User, err error) {
	var db = r.db.WithContext(ctx)

	condition := &entity.User{
		UserID: id,
	}

	query := db.Find(&result, condition)
	if query.Error != nil {
		r.log.Z().Err(query.Error).Msg("user.GetUserByID")
		return
	}

	return
}
