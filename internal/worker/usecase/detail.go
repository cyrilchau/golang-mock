package usecase

import (
	"context"
	"myapp/internal/worker/entity"
	"net/http"
)

func (uc *usecase) DetailWorker(ctx context.Context, id uint) (detail entity.Worker, httpCode int, err error) {
	workerDetail, err := uc.repo.GetWorkerByID(ctx, id)
	if err != nil {
		return detail, http.StatusInternalServerError, err
	}

	return workerDetail, http.StatusOK, nil
}
