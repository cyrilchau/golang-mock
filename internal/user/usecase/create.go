package usecase

import (
	"context"
	"net/http"

	"myapp/internal/user/dtos"
	"myapp/internal/user/entity"
	"myapp/pkg/errorCode"
)

func (uc *usecase) Create(ctx context.Context, payload dtos.CreateUserRequest) (result entity.User, httpCode int, err error) {
	user, err := uc.repo.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		uc.log.Z().Err(err).Msg("[usecase]CreateUser.GetUserByEmail")

		return result, http.StatusInternalServerError, err
	}

	if user.IsExists() {
		return result, http.StatusConflict, errorCode.ErrEmailAlreadyExist
	}

	result, err = uc.repo.CreateOneUser(ctx, dtos.NewCreateUser(payload, uc.cfg))
	if err != nil {
		uc.log.Z().Err(err).Msg("[usecase]CreateUser.SaveNewUser")

		return result, http.StatusInternalServerError, err
	}

	return result, http.StatusOK, nil
}
