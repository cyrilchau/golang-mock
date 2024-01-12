package delivery

import (
	"context"
	"myapp/internal/worker/dtos"
	"myapp/pkg/utils/response"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *handlers) Detail(c echo.Context) error {
	var (
		ctx, cancel = context.WithTimeout(c.Request().Context(), time.Duration(30*time.Second))
		payload     dtos.WorkerDetailRequest
	)
	defer cancel()

	data, code, err := h.uc.DetailWorker(ctx, payload.ID)
	if err != nil {
		return c.JSON(code, response.NewResponseError(code, response.MsgFailed, err.Error()))
	}

	return c.JSON(code, response.NewResponse(code, response.MsgSuccess, data))
}
