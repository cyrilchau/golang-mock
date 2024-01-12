package delivery

import (
	"myapp/internal/user/dtos"
	"myapp/pkg/utils/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handlers) Login(c echo.Context) error {
	var (
		request dtos.UserLoginRequest
	)

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(http.StatusBadRequest, response.MsgFailed, err.Error()))
	}

	if err := c.Validate(&request); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(http.StatusBadRequest, response.MsgFailed, err.Error()))
	}

	authData, httpCode, err := h.uc.Login(c.Request().Context(), request)
	if err != nil {
		return c.JSON(httpCode, response.NewResponseError(httpCode, response.MsgFailed, err.Error()))
	}

	return c.JSON(httpCode, response.NewResponse(httpCode, response.MsgSuccess, authData))
}
