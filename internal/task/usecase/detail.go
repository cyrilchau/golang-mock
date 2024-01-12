package usecase

import (
	"context"
	"net/http"

	"myapp/internal/task/dtos"
	"myapp/internal/task/entity"
)

func (uc *usecase) GetTaskDetail(ctx context.Context, payload dtos.DetailTaskRequest) (result entity.Task, httpCode int, err error) {
	result, err = uc.repo.GetTaskByID(ctx, payload.ID)
	if err != nil {
		uc.log.Z().Err(err).Msg("[usecase]Detail.GetTaskByID")

		return result, http.StatusInternalServerError, err
	}

	return result, http.StatusOK, nil
}
