package usecase

import (
	"context"
	"net/http"

	"myapp/internal/task/dtos"
	"myapp/internal/task/entity"
)

func (uc *usecase) GetListTask(ctx context.Context, payload dtos.ListTaskRequest) (result []entity.Task, httpCode int, err error) {
	result, err = uc.repo.GetAllTask(ctx, payload)
	if err != nil {
		uc.log.Z().Err(err).Msg("[usecase]List.GetAllTask")

		return result, http.StatusInternalServerError, err
	}

	return result, http.StatusOK, nil
}
