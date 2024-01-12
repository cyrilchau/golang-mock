package usecase

import (
	"context"
	"net/http"

	"myapp/internal/task/dtos"
	"myapp/internal/task/entity"
)

func (uc *usecase) CreateNewTask(ctx context.Context, payload dtos.CreateTaskRequest) (result entity.Task, httpCode int, err error) {
	result, err = uc.repo.CreateNewTask(ctx, dtos.NewCreateTask(payload, uc.cfg))
	if err != nil {
		uc.log.Z().Err(err).Msg("[usecase]Create.CreateNewTask")

		return result, http.StatusInternalServerError, err
	}

	return result, http.StatusOK, nil
}
