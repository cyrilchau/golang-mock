package usecase

import (
	"context"
	"myapp/internal/user/dtos"
	"net/http"
)

func (uc *usecase) Detail(ctx context.Context, id int) (detail dtos.UserDetailResponse, httpCode int, err error) {
	userDetail, err := uc.repo.GetUserByID(ctx, id)
	if err != nil {
		return detail, http.StatusInternalServerError, err
	}

	return dtos.NewUserDetail(userDetail), http.StatusOK, nil
}
