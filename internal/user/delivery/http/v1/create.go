package delivery

import (
	"context"
	"myapp/internal/user/dtos"
	"myapp/pkg/utils/response"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (s *service) Create(c echo.Context) error {
	var (
		ctx, cancel = context.WithTimeout(c.Request().Context(), time.Duration(30*time.Second))
		payload     dtos.CreateUserRequest
	)
	defer cancel()

	if err := c.Bind(&payload); err != nil {
		s.log.Z().Err(err).Msg("[service]CreateUser.Bind")

		return c.JSON(http.StatusBadRequest, response.NewResponseError(
			http.StatusBadRequest,
			response.MsgFailed,
			err.Error()),
		)
	}

	if err := c.Validate(&payload); err != nil {
		s.log.Z().Err(err).Msg("[service]CreateUser.Validate")

		return c.JSON(http.StatusBadRequest, response.NewResponseError(
			http.StatusBadRequest,
			response.MsgFailed,
			err.Error()),
		)
	}

	user, httpCode, err := s.uc.Create(ctx, payload)
	if err != nil {
		return c.JSON(httpCode, response.NewResponseError(
			httpCode,
			response.MsgFailed,
			err.Error()),
		)
	}

	res := dtos.CreateUserResponse{
		UserID: user.UserID,
		Fullname: user.Fullname,
		Email: user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	return c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, response.MsgSuccess, res))
}
