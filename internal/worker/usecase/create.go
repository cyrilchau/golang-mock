package usecase

import (
	"context"
	"net/http"

	"myapp/internal/worker/dtos"
	"myapp/internal/worker/entity"
)

func (uc *usecase) CreateOneWorker(ctx context.Context, payload dtos.CreateWorkerRequest) (result entity.Worker, httpCode int, err error) {
	result, err = uc.repo.CreateOneWorker(ctx, entity.NewCreateWorker(payload, uc.cfg))
	if err != nil {
		uc.log.Z().Err(err).Msg("[usecase]Create.CreateOneWorker")

		return result, http.StatusInternalServerError, err
	}

	return result, http.StatusOK, nil
}
