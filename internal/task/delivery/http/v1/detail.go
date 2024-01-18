package delivery

import (
	"context"
	"myapp/internal/task/dtos"
	"myapp/pkg/utils/response"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (s *service) Detail(c echo.Context) error {
	var (
		ctx, cancel = context.WithTimeout(c.Request().Context(), time.Duration(30*time.Second))
		payload     dtos.DetailTaskRequest
	)
	defer cancel()

	if err := c.Bind(&payload); err != nil {
		s.log.Z().Err(err).Msg("[service]DetailTask.Bind")

		return c.JSON(http.StatusBadRequest, response.NewResponseError(
			http.StatusBadRequest,
			response.MsgFailed,
			err.Error()),
		)
	}

	if err := c.Validate(&payload); err != nil {
		s.log.Z().Err(err).Msg("[service]DetailTask.Validate")

		return c.JSON(http.StatusBadRequest, response.NewResponseError(
			http.StatusBadRequest,
			response.MsgFailed,
			err.Error()),
		)
	}

	task, httpCode, err := s.uc.GetTaskDetail(ctx, payload)
	if err != nil {
		return c.JSON(httpCode, response.NewResponseError(
			httpCode,
			response.MsgFailed,
			err.Error()),
		)
	}

	return c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, response.MsgSuccess, task))
}
